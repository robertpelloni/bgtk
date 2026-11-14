#ifndef BOBGUI_CPP_JUCE_AUDIO_BRIDGE_HPP
#define BOBGUI_CPP_JUCE_AUDIO_BRIDGE_HPP

// Optional JUCE audio bridge (spike, see docs/JUCE_ANALYSIS.md).
//
// Why opt-in: JUCE is GPLv3/AGPLv3/commercial, Bobtk is LGPL. Linking this
// bridge changes distribution terms, so it compiles to nothing unless the
// embedder defines BOBGUI_WITH_JUCE and provides JUCE on the include path.
// Scope of this spike: device enumeration + sample-rate forwarding for
// media::AudioDeviceManager. processBlock forwarding is the next step and
// must wrap juce::AudioBuffer channels without copies.
//
// When enabled, JUCE modules required: juce_audio_devices, juce_audio_basics,
// juce_dsp (for JuceProcessorChain). NOTE: the BOBGUI_WITH_JUCE path is not
// compiled in CI on this machine (license + toolchain); it is reviewed code,
// syntax-verified only with the guard off. Verify before enabling.

#ifdef BOBGUI_WITH_JUCE

#include <juce_audio_devices/juce_audio_devices.h>
#include <juce_dsp/juce_dsp.h>

#include <string>
#include <vector>

namespace bobgui {
namespace cpp {
namespace media {

class JuceAudioBridge
{
public:
  JuceAudioBridge () = default;

  // Names of available audio device types (e.g. WASAPI, ASIO on Windows).
  std::vector<std::string> device_type_names ()
  {
    std::vector<std::string> names;
    for (auto *type : manager_.getAvailableDeviceTypes ())
      names.push_back (type->getTypeName ().toStdString ());
    return names;
  }

  // Current output device sample rate, or 0 when no device is open.
  double current_sample_rate () const
  {
    if (auto *device = manager_.getCurrentAudioDevice ())
      return device->getCurrentSampleRate ();
    return 0.0;
  }

  // Mirrors AudioDeviceManager::setSampleRate intent at the JUCE layer:
  // reopens the current setup with the requested rate when it differs.
  bool set_sample_rate (double sample_rate)
  {
    auto setup = manager_.getAudioDeviceSetup ();
    if (setup.sampleRate == sample_rate)
      return true;
    setup.sampleRate = sample_rate;
    return manager_.setAudioDeviceSetup (setup, true).isEmpty ();
  }

private:
  juce::AudioDeviceManager manager_;
};

// Forwards AudioProcessor::processBlock channel data through a JUCE DSP
// chain without copies: the float** channels are viewed as a
// juce::AudioBuffer and processed in place (ProcessContextReplacing).
class JuceProcessorChain
{
public:
  void prepare (double sample_rate, int max_block_size)
  {
    juce::dsp::ProcessSpec spec;
    spec.sampleRate = sample_rate;
    spec.maximumBlockSize = static_cast<juce::uint32> (max_block_size);
    spec.numChannels = 2;
    gain_.prepare (spec);
    gain_.setGainLinear (1.0f);
  }

  void set_gain (float linear_gain) { gain_.setGainLinear (linear_gain); }

  // Mirrors media::AudioProcessor::processBlock(float**, float**, int).
  // In-place processors (gain) use the input view directly; a future
  // non-replacing stage must copy in -> out first.
  void process_block (float **channels, int num_channels, int num_samples)
  {
    juce::AudioBuffer<float> buffer (channels, num_channels, num_samples);
    juce::dsp::AudioBlock<float> block (buffer);
    juce::dsp::ProcessContextReplacing<float> context (block);
    gain_.process (context);
  }

private:
  juce::dsp::Gain<float> gain_;
};

} /* namespace media */
} /* namespace cpp */
} /* namespace bobgui */

#endif /* BOBGUI_WITH_JUCE */

#endif /* BOBGUI_CPP_JUCE_AUDIO_BRIDGE_HPP */
