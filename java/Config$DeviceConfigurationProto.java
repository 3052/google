package com.google.android.gsf.checkin.proto;

import com.google.protobuf.micro.CodedInputStreamMicro;
import com.google.protobuf.micro.CodedOutputStreamMicro;
import com.google.protobuf.micro.MessageMicro;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
/* loaded from: classes.dex */
public final class Config$DeviceConfigurationProto extends MessageMicro {
    private boolean hasGlEsVersion;
    private boolean hasHasFiveWayNavigation;
    private boolean hasHasHardKeyboard;
    private boolean hasKeyboard;
    private boolean hasMaxApkDownloadSizeMb;
    private boolean hasNavigation;
    private boolean hasScreenDensity;
    private boolean hasScreenHeight;
    private boolean hasScreenLayout;
    private boolean hasScreenWidth;
    private boolean hasTouchScreen;
    private int touchScreen_ = 0;
    private int keyboard_ = 0;
    private int navigation_ = 0;
    private int screenLayout_ = 0;
    private boolean hasHardKeyboard_ = false;
    private boolean hasFiveWayNavigation_ = false;
    private int screenDensity_ = 0;
    private int screenWidth_ = 0;
    private int screenHeight_ = 0;
    private int glEsVersion_ = 0;
    private List<String> systemSharedLibrary_ = Collections.emptyList();
    private List<String> systemAvailableFeature_ = Collections.emptyList();
    private List<String> nativePlatform_ = Collections.emptyList();
    private List<String> systemSupportedLocale_ = Collections.emptyList();
    private List<String> glExtension_ = Collections.emptyList();
    private int maxApkDownloadSizeMb_ = 50;
    private int cachedSize = -1;

    public boolean hasTouchScreen() {
        return this.hasTouchScreen;
    }

    public int getTouchScreen() {
        return this.touchScreen_;
    }

    public Config$DeviceConfigurationProto setTouchScreen(int value) {
        this.hasTouchScreen = true;
        this.touchScreen_ = value;
        return this;
    }

    public boolean hasKeyboard() {
        return this.hasKeyboard;
    }

    public int getKeyboard() {
        return this.keyboard_;
    }

    public Config$DeviceConfigurationProto setKeyboard(int value) {
        this.hasKeyboard = true;
        this.keyboard_ = value;
        return this;
    }

    public boolean hasNavigation() {
        return this.hasNavigation;
    }

    public int getNavigation() {
        return this.navigation_;
    }

    public Config$DeviceConfigurationProto setNavigation(int value) {
        this.hasNavigation = true;
        this.navigation_ = value;
        return this;
    }

    public boolean hasScreenLayout() {
        return this.hasScreenLayout;
    }

    public int getScreenLayout() {
        return this.screenLayout_;
    }

    public Config$DeviceConfigurationProto setScreenLayout(int value) {
        this.hasScreenLayout = true;
        this.screenLayout_ = value;
        return this;
    }

    public boolean getHasHardKeyboard() {
        return this.hasHardKeyboard_;
    }

    public boolean hasHasHardKeyboard() {
        return this.hasHasHardKeyboard;
    }

    public Config$DeviceConfigurationProto setHasHardKeyboard(boolean value) {
        this.hasHasHardKeyboard = true;
        this.hasHardKeyboard_ = value;
        return this;
    }

    public boolean getHasFiveWayNavigation() {
        return this.hasFiveWayNavigation_;
    }

    public boolean hasHasFiveWayNavigation() {
        return this.hasHasFiveWayNavigation;
    }

    public Config$DeviceConfigurationProto setHasFiveWayNavigation(boolean value) {
        this.hasHasFiveWayNavigation = true;
        this.hasFiveWayNavigation_ = value;
        return this;
    }

    public int getScreenDensity() {
        return this.screenDensity_;
    }

    public boolean hasScreenDensity() {
        return this.hasScreenDensity;
    }

    public Config$DeviceConfigurationProto setScreenDensity(int value) {
        this.hasScreenDensity = true;
        this.screenDensity_ = value;
        return this;
    }

    public int getScreenWidth() {
        return this.screenWidth_;
    }

    public boolean hasScreenWidth() {
        return this.hasScreenWidth;
    }

    public Config$DeviceConfigurationProto setScreenWidth(int value) {
        this.hasScreenWidth = true;
        this.screenWidth_ = value;
        return this;
    }

    public int getScreenHeight() {
        return this.screenHeight_;
    }

    public boolean hasScreenHeight() {
        return this.hasScreenHeight;
    }

    public Config$DeviceConfigurationProto setScreenHeight(int value) {
        this.hasScreenHeight = true;
        this.screenHeight_ = value;
        return this;
    }

    public int getGlEsVersion() {
        return this.glEsVersion_;
    }

    public boolean hasGlEsVersion() {
        return this.hasGlEsVersion;
    }

    public Config$DeviceConfigurationProto setGlEsVersion(int value) {
        this.hasGlEsVersion = true;
        this.glEsVersion_ = value;
        return this;
    }

    public List<String> getSystemSharedLibraryList() {
        return this.systemSharedLibrary_;
    }

    public Config$DeviceConfigurationProto addSystemSharedLibrary(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.systemSharedLibrary_.isEmpty()) {
            this.systemSharedLibrary_ = new ArrayList();
        }
        this.systemSharedLibrary_.add(value);
        return this;
    }

    public List<String> getSystemAvailableFeatureList() {
        return this.systemAvailableFeature_;
    }

    public Config$DeviceConfigurationProto addSystemAvailableFeature(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.systemAvailableFeature_.isEmpty()) {
            this.systemAvailableFeature_ = new ArrayList();
        }
        this.systemAvailableFeature_.add(value);
        return this;
    }

    public List<String> getNativePlatformList() {
        return this.nativePlatform_;
    }

    public Config$DeviceConfigurationProto addNativePlatform(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.nativePlatform_.isEmpty()) {
            this.nativePlatform_ = new ArrayList();
        }
        this.nativePlatform_.add(value);
        return this;
    }

    public List<String> getSystemSupportedLocaleList() {
        return this.systemSupportedLocale_;
    }

    public Config$DeviceConfigurationProto addSystemSupportedLocale(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.systemSupportedLocale_.isEmpty()) {
            this.systemSupportedLocale_ = new ArrayList();
        }
        this.systemSupportedLocale_.add(value);
        return this;
    }

    public List<String> getGlExtensionList() {
        return this.glExtension_;
    }

    public Config$DeviceConfigurationProto addGlExtension(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.glExtension_.isEmpty()) {
            this.glExtension_ = new ArrayList();
        }
        this.glExtension_.add(value);
        return this;
    }

    public int getMaxApkDownloadSizeMb() {
        return this.maxApkDownloadSizeMb_;
    }

    public boolean hasMaxApkDownloadSizeMb() {
        return this.hasMaxApkDownloadSizeMb;
    }

    public Config$DeviceConfigurationProto setMaxApkDownloadSizeMb(int value) {
        this.hasMaxApkDownloadSizeMb = true;
        this.maxApkDownloadSizeMb_ = value;
        return this;
    }

    public final boolean isInitialized() {
        return this.hasTouchScreen && this.hasKeyboard && this.hasNavigation && this.hasScreenLayout && this.hasHasHardKeyboard && this.hasHasFiveWayNavigation && this.hasScreenDensity && this.hasGlEsVersion;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public void writeTo(CodedOutputStreamMicro output) throws IOException {
        if (hasTouchScreen()) {
            output.writeInt32(1, getTouchScreen());
        }
        if (hasKeyboard()) {
            output.writeInt32(2, getKeyboard());
        }
        if (hasNavigation()) {
            output.writeInt32(3, getNavigation());
        }
        if (hasScreenLayout()) {
            output.writeInt32(4, getScreenLayout());
        }
        if (hasHasHardKeyboard()) {
            output.writeBool(5, getHasHardKeyboard());
        }
        if (hasHasFiveWayNavigation()) {
            output.writeBool(6, getHasFiveWayNavigation());
        }
        if (hasScreenDensity()) {
            output.writeInt32(7, getScreenDensity());
        }
        if (hasGlEsVersion()) {
            output.writeInt32(8, getGlEsVersion());
        }
        for (String element : getSystemSharedLibraryList()) {
            output.writeString(9, element);
        }
        for (String element2 : getSystemAvailableFeatureList()) {
            output.writeString(10, element2);
        }
        for (String element3 : getNativePlatformList()) {
            output.writeString(11, element3);
        }
        if (hasScreenWidth()) {
            output.writeInt32(12, getScreenWidth());
        }
        if (hasScreenHeight()) {
            output.writeInt32(13, getScreenHeight());
        }
        for (String element4 : getSystemSupportedLocaleList()) {
            output.writeString(14, element4);
        }
        for (String element5 : getGlExtensionList()) {
            output.writeString(15, element5);
        }
        if (!hasMaxApkDownloadSizeMb()) {
            return;
        }
        output.writeInt32(17, getMaxApkDownloadSizeMb());
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public int getCachedSize() {
        if (this.cachedSize < 0) {
            getSerializedSize();
        }
        return this.cachedSize;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public int getSerializedSize() {
        int size = 0;
        if (hasTouchScreen()) {
            size = CodedOutputStreamMicro.computeInt32Size(1, getTouchScreen()) + 0;
        }
        if (hasKeyboard()) {
            size += CodedOutputStreamMicro.computeInt32Size(2, getKeyboard());
        }
        if (hasNavigation()) {
            size += CodedOutputStreamMicro.computeInt32Size(3, getNavigation());
        }
        if (hasScreenLayout()) {
            size += CodedOutputStreamMicro.computeInt32Size(4, getScreenLayout());
        }
        if (hasHasHardKeyboard()) {
            size += CodedOutputStreamMicro.computeBoolSize(5, getHasHardKeyboard());
        }
        if (hasHasFiveWayNavigation()) {
            size += CodedOutputStreamMicro.computeBoolSize(6, getHasFiveWayNavigation());
        }
        if (hasScreenDensity()) {
            size += CodedOutputStreamMicro.computeInt32Size(7, getScreenDensity());
        }
        if (hasGlEsVersion()) {
            size += CodedOutputStreamMicro.computeInt32Size(8, getGlEsVersion());
        }
        int dataSize = 0;
        for (String element : getSystemSharedLibraryList()) {
            dataSize += CodedOutputStreamMicro.computeStringSizeNoTag(element);
        }
        int size2 = size + dataSize + (getSystemSharedLibraryList().size() * 1);
        int dataSize2 = 0;
        for (String element2 : getSystemAvailableFeatureList()) {
            dataSize2 += CodedOutputStreamMicro.computeStringSizeNoTag(element2);
        }
        int size3 = size2 + dataSize2 + (getSystemAvailableFeatureList().size() * 1);
        int dataSize3 = 0;
        for (String element3 : getNativePlatformList()) {
            dataSize3 += CodedOutputStreamMicro.computeStringSizeNoTag(element3);
        }
        int size4 = size3 + dataSize3 + (getNativePlatformList().size() * 1);
        if (hasScreenWidth()) {
            size4 += CodedOutputStreamMicro.computeInt32Size(12, getScreenWidth());
        }
        if (hasScreenHeight()) {
            size4 += CodedOutputStreamMicro.computeInt32Size(13, getScreenHeight());
        }
        int dataSize4 = 0;
        for (String element4 : getSystemSupportedLocaleList()) {
            dataSize4 += CodedOutputStreamMicro.computeStringSizeNoTag(element4);
        }
        int size5 = size4 + dataSize4 + (getSystemSupportedLocaleList().size() * 1);
        int dataSize5 = 0;
        for (String element5 : getGlExtensionList()) {
            dataSize5 += CodedOutputStreamMicro.computeStringSizeNoTag(element5);
        }
        int size6 = size5 + dataSize5 + (getGlExtensionList().size() * 1);
        if (hasMaxApkDownloadSizeMb()) {
            size6 += CodedOutputStreamMicro.computeInt32Size(17, getMaxApkDownloadSizeMb());
        }
        this.cachedSize = size6;
        return size6;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public Config$DeviceConfigurationProto mergeFrom(CodedInputStreamMicro input) throws IOException {
        while (true) {
            int tag = input.readTag();
            switch (tag) {
                case 0:
                    return this;
                case 8:
                    setTouchScreen(input.readInt32());
                    break;
                case 16:
                    setKeyboard(input.readInt32());
                    break;
                case 24:
                    setNavigation(input.readInt32());
                    break;
                case 32:
                    setScreenLayout(input.readInt32());
                    break;
                case 40:
                    setHasHardKeyboard(input.readBool());
                    break;
                case 48:
                    setHasFiveWayNavigation(input.readBool());
                    break;
                case 56:
                    setScreenDensity(input.readInt32());
                    break;
                case 64:
                    setGlEsVersion(input.readInt32());
                    break;
                case 74:
                    addSystemSharedLibrary(input.readString());
                    break;
                case 82:
                    addSystemAvailableFeature(input.readString());
                    break;
                case 90:
                    addNativePlatform(input.readString());
                    break;
                case 96:
                    setScreenWidth(input.readInt32());
                    break;
                case 104:
                    setScreenHeight(input.readInt32());
                    break;
                case 114:
                    addSystemSupportedLocale(input.readString());
                    break;
                case 122:
                    addGlExtension(input.readString());
                    break;
                case 136:
                    setMaxApkDownloadSizeMb(input.readInt32());
                    break;
                default:
                    if (parseUnknownField(input, tag)) {
                        break;
                    } else {
                        return this;
                    }
            }
        }
    }
}
