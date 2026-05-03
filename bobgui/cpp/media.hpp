#pragma once

#include <bobgui/bobgui.h>
#include <bobgui/cpp/object_handle.hpp>
#include <bobgui/cpp/widget.hpp>
#include <string>
#include <memory>
#include <vector>

namespace bobtk {
namespace media {

// ---------------------------------------------------------
// AUDIO SUBSYSTEM (Parity with QtMultimedia / JUCE)
// ---------------------------------------------------------

class AudioDeviceManager : public core::ObjectHandle<BobguiAudioDeviceManager> {
public:
    AudioDeviceManager()
        : core::ObjectHandle<BobguiAudioDeviceManager>(bobgui_audio_device_manager_new()) {}

    void initializeDevices(int inputChannels, int outputChannels) {
        bobgui_audio_device_manager_init_devices(handle(), inputChannels, outputChannels);
    }

    void setSampleRate(double sampleRate) {
        bobgui_audio_device_manager_set_sample_rate(handle(), sampleRate);
    }

    void play() {
        bobgui_audio_device_manager_play(handle());
    }

    void stop() {
        bobgui_audio_device_manager_stop(handle());
    }
};

class AudioProcessor : public core::ObjectHandle<BobguiAudioProcessor> {
public:
    explicit AudioProcessor(const std::string& name)
        : core::ObjectHandle<BobguiAudioProcessor>(bobgui_audio_processor_new(name.c_str())) {}

    void setParameter(int index, float value) {
        bobgui_audio_processor_set_parameter(handle(), index, value);
    }

    float getParameter(int index) const {
        return bobgui_audio_processor_get_parameter(handle(), index);
    }

    void processBlock(float** inChannels, float** outChannels, int numSamples) {
        bobgui_audio_processor_process_block(handle(), inChannels, outChannels, numSamples);
    }
};

class AudioDial : public Widget {
public:
    AudioDial() : Widget(bobgui_audio_dial_new()) {}
};

class AudioFader : public Widget {
public:
    AudioFader() : Widget(bobgui_audio_fader_new()) {}
};

class AudioScope : public Widget {
public:
    AudioScope() : Widget(bobgui_audio_scope_new()) {}
};

class AudioAnalyzer : public Widget {
public:
    AudioAnalyzer() : Widget(bobgui_audio_analyzer_new()) {}
};

// ---------------------------------------------------------
// 3D SUBSYSTEM (Parity with Qt3D)
// ---------------------------------------------------------

class ThreeDNode : public core::ObjectHandle<Bobgui3dNode> {
public:
    ThreeDNode() : core::ObjectHandle<Bobgui3dNode>(bobgui_3d_node_new()) {}

    void setTranslation(float x, float y, float z) {
        bobgui_3d_node_set_translation(handle(), x, y, z);
    }

    void setRotation(float pitch, float yaw, float roll) {
        bobgui_3d_node_set_rotation(handle(), pitch, yaw, roll);
    }

    void setScale(float scale) {
        bobgui_3d_node_set_scale(handle(), scale);
    }

    void loadMesh(const std::string& path) {
        bobgui_3d_node_load_mesh(handle(), path.c_str());
    }
};

// ---------------------------------------------------------
// HOLOGRAPHIC SUBSYSTEM
// ---------------------------------------------------------

class Hologram : public core::ObjectHandle<BobguiHologram> {
public:
    Hologram(gpointer voxels, int width, int height, int depth)
        : core::ObjectHandle<BobguiHologram>(bobgui_hologram_new_from_voxels(voxels, width, height, depth)) {}

    void projectToSpace(float x, float y, float z, float w, float h, float d) {
        bobgui_hologram_project_to_space(handle(), x, y, z, w, h, d);
    }

    void setInteractionDensity(float threshold) {
        bobgui_hologram_set_interaction_density(handle(), threshold);
    }
};

class HolographView : public Widget {
public:
    HolographView() : Widget(bobgui_holograph_view_new()) {}

    void display(const Hologram& hologram) {
        bobgui_holograph_view_display(BOBGUI_HOLOGRAPH_VIEW(handle()), hologram.handle());
    }
};

// ---------------------------------------------------------
// GIS / MAP SUBSYSTEM (Parity with QtLocation)
// ---------------------------------------------------------

class MapView : public Widget {
public:
    MapView() : Widget(bobgui_map_view_new()) {}

    void setCenter(double latitude, double longitude) {
        bobgui_map_view_set_center(BOBGUI_MAP_VIEW(handle()), latitude, longitude);
    }

    void setZoom(double zoomLevel) {
        bobgui_map_view_set_zoom(BOBGUI_MAP_VIEW(handle()), zoomLevel);
    }

    void addMarker(double latitude, double longitude, const std::string& title) {
        bobgui_map_view_add_marker(BOBGUI_MAP_VIEW(handle()), latitude, longitude, title.c_str());
    }
};

// ---------------------------------------------------------
// SHADER SUBSYSTEM
// ---------------------------------------------------------

class EffectShader : public core::ObjectHandle<BobguiEffectShader> {
public:
    explicit EffectShader(const std::string& source)
        : core::ObjectHandle<BobguiEffectShader>(bobgui_effect_shader_new_from_source(source.c_str())) {}

    void setUniformFloat(const std::string& name, float value) {
        bobgui_effect_shader_set_uniform_float(handle(), name.c_str(), value);
    }

    void setUniformVec3(const std::string& name, float x, float y, float z) {
        bobgui_effect_shader_set_uniform_vec3(handle(), name.c_str(), x, y, z);
    }
};

// ---------------------------------------------------------
// PHYSICS SUBSYSTEM
// ---------------------------------------------------------

class PhysicsWorld : public core::ObjectHandle<BobguiPhysicsWorld> {
public:
    PhysicsWorld() : core::ObjectHandle<BobguiPhysicsWorld>(bobgui_physics_world_new()) {}

    void step(float deltaTime) {
        bobgui_physics_world_step(handle(), deltaTime);
    }

    void setGravity(float x, float y, float z) {
        bobgui_physics_world_set_gravity(handle(), x, y, z);
    }
};

// ---------------------------------------------------------
// TIMELINE SUBSYSTEM
// ---------------------------------------------------------

class Timeline : public core::ObjectHandle<BobguiTimeline> {
public:
    Timeline() : core::ObjectHandle<BobguiTimeline>(bobgui_timeline_new()) {}

    void play() {
        bobgui_timeline_play(handle());
    }

    void pause() {
        bobgui_timeline_pause(handle());
    }

    void seek(double timeSeconds) {
        bobgui_timeline_seek(handle(), timeSeconds);
    }
};

// ---------------------------------------------------------
// BIO SUBSYSTEM
// ---------------------------------------------------------

class BioManager : public core::ObjectHandle<BobguiBioManager> {
public:
    static std::shared_ptr<BioManager> getDefault() {
        return std::shared_ptr<BioManager>(new BioManager(bobgui_bio_manager_get_default()));
    }

    void startSensorStream(const std::string& sensorId) {
        bobgui_bio_start_sensor_stream(handle(), sensorId.c_str());
    }

    void enableBiometricAdaptation(Widget& uiRoot) {
        bobgui_bio_enable_biometric_adaptation(handle(), BOBGUI_WIDGET(uiRoot.handle()));
    }

private:
    BioManager(BobguiBioManager* handle) : core::ObjectHandle<BobguiBioManager>(handle) {}
};

class BioViewer : public Widget {
public:
    BioViewer() : Widget(bobgui_bio_viewer_new()) {}

    void loadDicom(const std::string& path) {
        bobgui_bio_viewer_load_dicom(BOBGUI_BIO_VIEWER(handle()), path.c_str());
    }
};

// ---------------------------------------------------------
// SPATIAL SUBSYSTEM
// ---------------------------------------------------------

class SpatialContext : public core::ObjectHandle<BobguiSpatialContext> {
public:
    explicit SpatialContext(const std::string& runtimeName)
        : core::ObjectHandle<BobguiSpatialContext>(bobgui_spatial_context_new(runtimeName.c_str())) {}
};

} // namespace media
} // namespace bobtk
