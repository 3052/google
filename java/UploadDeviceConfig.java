package com.google.android.finsky.protos;

import com.google.android.finsky.protos.DeviceConfiguration;
import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public interface UploadDeviceConfig {

    /* loaded from: classes.dex */
    public static final class UploadDeviceConfigRequest extends MessageNano {
        public DeviceConfiguration.DeviceConfigurationProto deviceConfiguration = null;
        public String manufacturer = "";
        public boolean hasManufacturer = false;
        public DataServiceSubscriber dataServiceSubscriber = null;
        public String gcmRegistrationId = "";
        public boolean hasGcmRegistrationId = false;

        public UploadDeviceConfigRequest() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.deviceConfiguration != null) {
                output.writeMessage(1, this.deviceConfiguration);
            }
            if (this.hasManufacturer || !this.manufacturer.equals("")) {
                output.writeString(2, this.manufacturer);
            }
            if (this.hasGcmRegistrationId || !this.gcmRegistrationId.equals("")) {
                output.writeString(3, this.gcmRegistrationId);
            }
            if (this.dataServiceSubscriber != null) {
                output.writeMessage(4, this.dataServiceSubscriber);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.deviceConfiguration != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.deviceConfiguration);
            }
            if (this.hasManufacturer || !this.manufacturer.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.manufacturer);
            }
            if (this.hasGcmRegistrationId || !this.gcmRegistrationId.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(3, this.gcmRegistrationId);
            }
            if (this.dataServiceSubscriber != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(4, this.dataServiceSubscriber);
            }
            return size;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        if (this.deviceConfiguration == null) {
                            this.deviceConfiguration = new DeviceConfiguration.DeviceConfigurationProto();
                        }
                        x0.readMessage(this.deviceConfiguration);
                        break;
                    case 18:
                        this.manufacturer = x0.readString();
                        this.hasManufacturer = true;
                        break;
                    case 26:
                        this.gcmRegistrationId = x0.readString();
                        this.hasGcmRegistrationId = true;
                        break;
                    case 34:
                        if (this.dataServiceSubscriber == null) {
                            this.dataServiceSubscriber = new DataServiceSubscriber();
                        }
                        x0.readMessage(this.dataServiceSubscriber);
                        break;
                    default:
                        if (!WireFormatNano.parseUnknownField(x0, readTag)) {
                            break;
                        } else {
                            break;
                        }
                }
            }
            return this;
        }
    }

    /* loaded from: classes.dex */
    public static final class UploadDeviceConfigResponse extends MessageNano {
        public String uploadDeviceConfigToken = "";
        public boolean hasUploadDeviceConfigToken = false;

        public UploadDeviceConfigResponse() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasUploadDeviceConfigToken || !this.uploadDeviceConfigToken.equals("")) {
                output.writeString(1, this.uploadDeviceConfigToken);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasUploadDeviceConfigToken || !this.uploadDeviceConfigToken.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(1, this.uploadDeviceConfigToken);
            }
            return size;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.uploadDeviceConfigToken = x0.readString();
                        this.hasUploadDeviceConfigToken = true;
                        break;
                    default:
                        if (!WireFormatNano.parseUnknownField(x0, readTag)) {
                            break;
                        } else {
                            break;
                        }
                }
            }
            return this;
        }
    }
}
