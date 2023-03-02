package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class AndroidAppDeliveryData extends MessageNano {
    public long downloadSize = 0;
    public boolean hasDownloadSize = false;
    public long gzippedDownloadSize = 0;
    public boolean hasGzippedDownloadSize = false;
    public String signature = "";
    public boolean hasSignature = false;
    public String downloadUrl = "";
    public boolean hasDownloadUrl = false;
    public String gzippedDownloadUrl = "";
    public boolean hasGzippedDownloadUrl = false;
    public EncryptionParams encryptionParams = null;
    public AppFileMetadata[] additionalFile = AppFileMetadata.emptyArray();
    public HttpCookie[] downloadAuthCookie = HttpCookie.emptyArray();
    public boolean forwardLocked = false;
    public boolean hasForwardLocked = false;
    public long refundTimeout = 0;
    public boolean hasRefundTimeout = false;
    public long postInstallRefundWindowMillis = 0;
    public boolean hasPostInstallRefundWindowMillis = false;
    public boolean serverInitiated = true;
    public boolean hasServerInitiated = false;
    public boolean immediateStartNeeded = false;
    public boolean hasImmediateStartNeeded = false;
    public AndroidAppPatchData patchData = null;
    public SplitDeliveryData[] splitDeliveryData = SplitDeliveryData.emptyArray();
    public int installLocation = 0;
    public boolean hasInstallLocation = false;
    public boolean everExternallyHosted = false;
    public boolean hasEverExternallyHosted = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
        int length2;
        int length3;
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 8:
                    this.downloadSize = x0.readRawVarint64();
                    this.hasDownloadSize = true;
                    break;
                case 18:
                    this.signature = x0.readString();
                    this.hasSignature = true;
                    break;
                case 26:
                    this.downloadUrl = x0.readString();
                    this.hasDownloadUrl = true;
                    break;
                case 34:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 34);
                    if (this.additionalFile == null) {
                        length3 = 0;
                    } else {
                        length3 = this.additionalFile.length;
                    }
                    AppFileMetadata[] appFileMetadataArr = new AppFileMetadata[repeatedFieldArrayLength + length3];
                    if (length3 != 0) {
                        System.arraycopy(this.additionalFile, 0, appFileMetadataArr, 0, length3);
                    }
                    while (length3 < appFileMetadataArr.length - 1) {
                        appFileMetadataArr[length3] = new AppFileMetadata();
                        x0.readMessage(appFileMetadataArr[length3]);
                        x0.readTag();
                        length3++;
                    }
                    appFileMetadataArr[length3] = new AppFileMetadata();
                    x0.readMessage(appFileMetadataArr[length3]);
                    this.additionalFile = appFileMetadataArr;
                    break;
                case 42:
                    int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 42);
                    if (this.downloadAuthCookie == null) {
                        length2 = 0;
                    } else {
                        length2 = this.downloadAuthCookie.length;
                    }
                    HttpCookie[] httpCookieArr = new HttpCookie[repeatedFieldArrayLength2 + length2];
                    if (length2 != 0) {
                        System.arraycopy(this.downloadAuthCookie, 0, httpCookieArr, 0, length2);
                    }
                    while (length2 < httpCookieArr.length - 1) {
                        httpCookieArr[length2] = new HttpCookie();
                        x0.readMessage(httpCookieArr[length2]);
                        x0.readTag();
                        length2++;
                    }
                    httpCookieArr[length2] = new HttpCookie();
                    x0.readMessage(httpCookieArr[length2]);
                    this.downloadAuthCookie = httpCookieArr;
                    break;
                case 48:
                    this.forwardLocked = x0.readBool();
                    this.hasForwardLocked = true;
                    break;
                case 56:
                    this.refundTimeout = x0.readRawVarint64();
                    this.hasRefundTimeout = true;
                    break;
                case 64:
                    this.serverInitiated = x0.readBool();
                    this.hasServerInitiated = true;
                    break;
                case 72:
                    this.postInstallRefundWindowMillis = x0.readRawVarint64();
                    this.hasPostInstallRefundWindowMillis = true;
                    break;
                case 80:
                    this.immediateStartNeeded = x0.readBool();
                    this.hasImmediateStartNeeded = true;
                    break;
                case 90:
                    if (this.patchData == null) {
                        this.patchData = new AndroidAppPatchData();
                    }
                    x0.readMessage(this.patchData);
                    break;
                case 98:
                    if (this.encryptionParams == null) {
                        this.encryptionParams = new EncryptionParams();
                    }
                    x0.readMessage(this.encryptionParams);
                    break;
                case 106:
                    this.gzippedDownloadUrl = x0.readString();
                    this.hasGzippedDownloadUrl = true;
                    break;
                case 112:
                    this.gzippedDownloadSize = x0.readRawVarint64();
                    this.hasGzippedDownloadSize = true;
                    break;
                case 122:
                    int repeatedFieldArrayLength3 = WireFormatNano.getRepeatedFieldArrayLength(x0, 122);
                    if (this.splitDeliveryData == null) {
                        length = 0;
                    } else {
                        length = this.splitDeliveryData.length;
                    }
                    SplitDeliveryData[] splitDeliveryDataArr = new SplitDeliveryData[repeatedFieldArrayLength3 + length];
                    if (length != 0) {
                        System.arraycopy(this.splitDeliveryData, 0, splitDeliveryDataArr, 0, length);
                    }
                    while (length < splitDeliveryDataArr.length - 1) {
                        splitDeliveryDataArr[length] = new SplitDeliveryData();
                        x0.readMessage(splitDeliveryDataArr[length]);
                        x0.readTag();
                        length++;
                    }
                    splitDeliveryDataArr[length] = new SplitDeliveryData();
                    x0.readMessage(splitDeliveryDataArr[length]);
                    this.splitDeliveryData = splitDeliveryDataArr;
                    break;
                case 128:
                    int readRawVarint32 = x0.readRawVarint32();
                    switch (readRawVarint32) {
                        case 0:
                        case 1:
                        case 2:
                        case 3:
                            this.installLocation = readRawVarint32;
                            this.hasInstallLocation = true;
                            continue;
                    }
                case 136:
                    this.everExternallyHosted = x0.readBool();
                    this.hasEverExternallyHosted = true;
                    break;
                default:
                    if (WireFormatNano.parseUnknownField(x0, readTag)) {
                        break;
                    } else {
                        break;
                    }
            }
        }
        return this;
    }

    public AndroidAppDeliveryData() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasDownloadSize || this.downloadSize != 0) {
            output.writeInt64(1, this.downloadSize);
        }
        if (this.hasSignature || !this.signature.equals("")) {
            output.writeString(2, this.signature);
        }
        if (this.hasDownloadUrl || !this.downloadUrl.equals("")) {
            output.writeString(3, this.downloadUrl);
        }
        if (this.additionalFile != null && this.additionalFile.length > 0) {
            for (int i = 0; i < this.additionalFile.length; i++) {
                AppFileMetadata element = this.additionalFile[i];
                if (element != null) {
                    output.writeMessage(4, element);
                }
            }
        }
        if (this.downloadAuthCookie != null && this.downloadAuthCookie.length > 0) {
            for (int i2 = 0; i2 < this.downloadAuthCookie.length; i2++) {
                HttpCookie element2 = this.downloadAuthCookie[i2];
                if (element2 != null) {
                    output.writeMessage(5, element2);
                }
            }
        }
        if (this.hasForwardLocked || this.forwardLocked) {
            output.writeBool(6, this.forwardLocked);
        }
        if (this.hasRefundTimeout || this.refundTimeout != 0) {
            output.writeInt64(7, this.refundTimeout);
        }
        if (this.hasServerInitiated || !this.serverInitiated) {
            output.writeBool(8, this.serverInitiated);
        }
        if (this.hasPostInstallRefundWindowMillis || this.postInstallRefundWindowMillis != 0) {
            output.writeInt64(9, this.postInstallRefundWindowMillis);
        }
        if (this.hasImmediateStartNeeded || this.immediateStartNeeded) {
            output.writeBool(10, this.immediateStartNeeded);
        }
        if (this.patchData != null) {
            output.writeMessage(11, this.patchData);
        }
        if (this.encryptionParams != null) {
            output.writeMessage(12, this.encryptionParams);
        }
        if (this.hasGzippedDownloadUrl || !this.gzippedDownloadUrl.equals("")) {
            output.writeString(13, this.gzippedDownloadUrl);
        }
        if (this.hasGzippedDownloadSize || this.gzippedDownloadSize != 0) {
            output.writeInt64(14, this.gzippedDownloadSize);
        }
        if (this.splitDeliveryData != null && this.splitDeliveryData.length > 0) {
            for (int i3 = 0; i3 < this.splitDeliveryData.length; i3++) {
                SplitDeliveryData element3 = this.splitDeliveryData[i3];
                if (element3 != null) {
                    output.writeMessage(15, element3);
                }
            }
        }
        if (this.installLocation != 0 || this.hasInstallLocation) {
            output.writeInt32(16, this.installLocation);
        }
        if (this.hasEverExternallyHosted || this.everExternallyHosted) {
            output.writeBool(17, this.everExternallyHosted);
        }
        super.writeTo(output);
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasDownloadSize || this.downloadSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(1, this.downloadSize);
        }
        if (this.hasSignature || !this.signature.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(2, this.signature);
        }
        if (this.hasDownloadUrl || !this.downloadUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(3, this.downloadUrl);
        }
        if (this.additionalFile != null && this.additionalFile.length > 0) {
            for (int i = 0; i < this.additionalFile.length; i++) {
                AppFileMetadata element = this.additionalFile[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(4, element);
                }
            }
        }
        if (this.downloadAuthCookie != null && this.downloadAuthCookie.length > 0) {
            for (int i2 = 0; i2 < this.downloadAuthCookie.length; i2++) {
                HttpCookie element2 = this.downloadAuthCookie[i2];
                if (element2 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(5, element2);
                }
            }
        }
        if (this.hasForwardLocked || this.forwardLocked) {
            size += CodedOutputByteBufferNano.computeTagSize(6) + 1;
        }
        if (this.hasRefundTimeout || this.refundTimeout != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(7, this.refundTimeout);
        }
        if (this.hasServerInitiated || !this.serverInitiated) {
            size += CodedOutputByteBufferNano.computeTagSize(8) + 1;
        }
        if (this.hasPostInstallRefundWindowMillis || this.postInstallRefundWindowMillis != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(9, this.postInstallRefundWindowMillis);
        }
        if (this.hasImmediateStartNeeded || this.immediateStartNeeded) {
            size += CodedOutputByteBufferNano.computeTagSize(10) + 1;
        }
        if (this.patchData != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(11, this.patchData);
        }
        if (this.encryptionParams != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(12, this.encryptionParams);
        }
        if (this.hasGzippedDownloadUrl || !this.gzippedDownloadUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(13, this.gzippedDownloadUrl);
        }
        if (this.hasGzippedDownloadSize || this.gzippedDownloadSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(14, this.gzippedDownloadSize);
        }
        if (this.splitDeliveryData != null && this.splitDeliveryData.length > 0) {
            for (int i3 = 0; i3 < this.splitDeliveryData.length; i3++) {
                SplitDeliveryData element3 = this.splitDeliveryData[i3];
                if (element3 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(15, element3);
                }
            }
        }
        if (this.installLocation != 0 || this.hasInstallLocation) {
            size += CodedOutputByteBufferNano.computeInt32Size(16, this.installLocation);
        }
        if (this.hasEverExternallyHosted || this.everExternallyHosted) {
            return size + CodedOutputByteBufferNano.computeTagSize(17) + 1;
        }
        return size;
    }
}
