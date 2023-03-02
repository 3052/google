package com.google.android.gsf.checkin.proto;

import com.google.protobuf.micro.CodedInputStreamMicro;
import com.google.protobuf.micro.CodedOutputStreamMicro;
import com.google.protobuf.micro.MessageMicro;
import java.io.IOException;
/* loaded from: classes.dex */
public final class Logs$AndroidBuildProto extends MessageMicro {
    private boolean hasBootloader;
    private boolean hasBuildProduct;
    private boolean hasCarrier;
    private boolean hasClient;
    private boolean hasDevice;
    private boolean hasGoogleServices;
    private boolean hasId;
    private boolean hasManufacturer;
    private boolean hasModel;
    private boolean hasOtaInstalled;
    private boolean hasProduct;
    private boolean hasRadio;
    private boolean hasSdkVersion;
    private boolean hasTimestamp;
    private String id_ = "";
    private String radio_ = "";
    private String bootloader_ = "";
    private String product_ = "";
    private String carrier_ = "";
    private String client_ = "";
    private long timestamp_ = 0;
    private int googleServices_ = 0;
    private int sdkVersion_ = 0;
    private String device_ = "";
    private String model_ = "";
    private String manufacturer_ = "";
    private String buildProduct_ = "";
    private boolean otaInstalled_ = false;
    private int cachedSize = -1;

    public String getId() {
        return this.id_;
    }

    public boolean hasId() {
        return this.hasId;
    }

    public Logs$AndroidBuildProto setId(String value) {
        this.hasId = true;
        this.id_ = value;
        return this;
    }

    public String getRadio() {
        return this.radio_;
    }

    public boolean hasRadio() {
        return this.hasRadio;
    }

    public Logs$AndroidBuildProto setRadio(String value) {
        this.hasRadio = true;
        this.radio_ = value;
        return this;
    }

    public String getBootloader() {
        return this.bootloader_;
    }

    public boolean hasBootloader() {
        return this.hasBootloader;
    }

    public Logs$AndroidBuildProto setBootloader(String value) {
        this.hasBootloader = true;
        this.bootloader_ = value;
        return this;
    }

    public String getProduct() {
        return this.product_;
    }

    public boolean hasProduct() {
        return this.hasProduct;
    }

    public Logs$AndroidBuildProto setProduct(String value) {
        this.hasProduct = true;
        this.product_ = value;
        return this;
    }

    public String getCarrier() {
        return this.carrier_;
    }

    public boolean hasCarrier() {
        return this.hasCarrier;
    }

    public Logs$AndroidBuildProto setCarrier(String value) {
        this.hasCarrier = true;
        this.carrier_ = value;
        return this;
    }

    public String getClient() {
        return this.client_;
    }

    public boolean hasClient() {
        return this.hasClient;
    }

    public Logs$AndroidBuildProto setClient(String value) {
        this.hasClient = true;
        this.client_ = value;
        return this;
    }

    public long getTimestamp() {
        return this.timestamp_;
    }

    public boolean hasTimestamp() {
        return this.hasTimestamp;
    }

    public Logs$AndroidBuildProto setTimestamp(long value) {
        this.hasTimestamp = true;
        this.timestamp_ = value;
        return this;
    }

    public int getGoogleServices() {
        return this.googleServices_;
    }

    public boolean hasGoogleServices() {
        return this.hasGoogleServices;
    }

    public Logs$AndroidBuildProto setGoogleServices(int value) {
        this.hasGoogleServices = true;
        this.googleServices_ = value;
        return this;
    }

    public int getSdkVersion() {
        return this.sdkVersion_;
    }

    public boolean hasSdkVersion() {
        return this.hasSdkVersion;
    }

    public Logs$AndroidBuildProto setSdkVersion(int value) {
        this.hasSdkVersion = true;
        this.sdkVersion_ = value;
        return this;
    }

    public String getDevice() {
        return this.device_;
    }

    public boolean hasDevice() {
        return this.hasDevice;
    }

    public Logs$AndroidBuildProto setDevice(String value) {
        this.hasDevice = true;
        this.device_ = value;
        return this;
    }

    public String getModel() {
        return this.model_;
    }

    public boolean hasModel() {
        return this.hasModel;
    }

    public Logs$AndroidBuildProto setModel(String value) {
        this.hasModel = true;
        this.model_ = value;
        return this;
    }

    public String getManufacturer() {
        return this.manufacturer_;
    }

    public boolean hasManufacturer() {
        return this.hasManufacturer;
    }

    public Logs$AndroidBuildProto setManufacturer(String value) {
        this.hasManufacturer = true;
        this.manufacturer_ = value;
        return this;
    }

    public String getBuildProduct() {
        return this.buildProduct_;
    }

    public boolean hasBuildProduct() {
        return this.hasBuildProduct;
    }

    public Logs$AndroidBuildProto setBuildProduct(String value) {
        this.hasBuildProduct = true;
        this.buildProduct_ = value;
        return this;
    }

    public boolean getOtaInstalled() {
        return this.otaInstalled_;
    }

    public boolean hasOtaInstalled() {
        return this.hasOtaInstalled;
    }

    public Logs$AndroidBuildProto setOtaInstalled(boolean value) {
        this.hasOtaInstalled = true;
        this.otaInstalled_ = value;
        return this;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public void writeTo(CodedOutputStreamMicro output) throws IOException {
        if (hasId()) {
            output.writeString(1, getId());
        }
        if (hasProduct()) {
            output.writeString(2, getProduct());
        }
        if (hasCarrier()) {
            output.writeString(3, getCarrier());
        }
        if (hasRadio()) {
            output.writeString(4, getRadio());
        }
        if (hasBootloader()) {
            output.writeString(5, getBootloader());
        }
        if (hasClient()) {
            output.writeString(6, getClient());
        }
        if (hasTimestamp()) {
            output.writeInt64(7, getTimestamp());
        }
        if (hasGoogleServices()) {
            output.writeInt32(8, getGoogleServices());
        }
        if (hasDevice()) {
            output.writeString(9, getDevice());
        }
        if (hasSdkVersion()) {
            output.writeInt32(10, getSdkVersion());
        }
        if (hasModel()) {
            output.writeString(11, getModel());
        }
        if (hasManufacturer()) {
            output.writeString(12, getManufacturer());
        }
        if (hasBuildProduct()) {
            output.writeString(13, getBuildProduct());
        }
        if (!hasOtaInstalled()) {
            return;
        }
        output.writeBool(14, getOtaInstalled());
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
        if (hasId()) {
            size = CodedOutputStreamMicro.computeStringSize(1, getId()) + 0;
        }
        if (hasProduct()) {
            size += CodedOutputStreamMicro.computeStringSize(2, getProduct());
        }
        if (hasCarrier()) {
            size += CodedOutputStreamMicro.computeStringSize(3, getCarrier());
        }
        if (hasRadio()) {
            size += CodedOutputStreamMicro.computeStringSize(4, getRadio());
        }
        if (hasBootloader()) {
            size += CodedOutputStreamMicro.computeStringSize(5, getBootloader());
        }
        if (hasClient()) {
            size += CodedOutputStreamMicro.computeStringSize(6, getClient());
        }
        if (hasTimestamp()) {
            size += CodedOutputStreamMicro.computeInt64Size(7, getTimestamp());
        }
        if (hasGoogleServices()) {
            size += CodedOutputStreamMicro.computeInt32Size(8, getGoogleServices());
        }
        if (hasDevice()) {
            size += CodedOutputStreamMicro.computeStringSize(9, getDevice());
        }
        if (hasSdkVersion()) {
            size += CodedOutputStreamMicro.computeInt32Size(10, getSdkVersion());
        }
        if (hasModel()) {
            size += CodedOutputStreamMicro.computeStringSize(11, getModel());
        }
        if (hasManufacturer()) {
            size += CodedOutputStreamMicro.computeStringSize(12, getManufacturer());
        }
        if (hasBuildProduct()) {
            size += CodedOutputStreamMicro.computeStringSize(13, getBuildProduct());
        }
        if (hasOtaInstalled()) {
            size += CodedOutputStreamMicro.computeBoolSize(14, getOtaInstalled());
        }
        this.cachedSize = size;
        return size;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public Logs$AndroidBuildProto mergeFrom(CodedInputStreamMicro input) throws IOException {
        while (true) {
            int tag = input.readTag();
            switch (tag) {
                case 0:
                    return this;
                case 10:
                    setId(input.readString());
                    break;
                case 18:
                    setProduct(input.readString());
                    break;
                case 26:
                    setCarrier(input.readString());
                    break;
                case 34:
                    setRadio(input.readString());
                    break;
                case 42:
                    setBootloader(input.readString());
                    break;
                case 50:
                    setClient(input.readString());
                    break;
                case 56:
                    setTimestamp(input.readInt64());
                    break;
                case 64:
                    setGoogleServices(input.readInt32());
                    break;
                case 74:
                    setDevice(input.readString());
                    break;
                case 80:
                    setSdkVersion(input.readInt32());
                    break;
                case 90:
                    setModel(input.readString());
                    break;
                case 98:
                    setManufacturer(input.readString());
                    break;
                case 106:
                    setBuildProduct(input.readString());
                    break;
                case 112:
                    setOtaInstalled(input.readBool());
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
