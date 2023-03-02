package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class SplitDeliveryData extends MessageNano {
    private static volatile SplitDeliveryData[] _emptyArray;
    public String id = "";
    public boolean hasId = false;
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
    public AndroidAppPatchData patchData = null;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    this.id = x0.readString();
                    this.hasId = true;
                    break;
                case 16:
                    this.downloadSize = x0.readRawVarint64();
                    this.hasDownloadSize = true;
                    break;
                case 24:
                    this.gzippedDownloadSize = x0.readRawVarint64();
                    this.hasGzippedDownloadSize = true;
                    break;
                case 34:
                    this.signature = x0.readString();
                    this.hasSignature = true;
                    break;
                case 42:
                    this.downloadUrl = x0.readString();
                    this.hasDownloadUrl = true;
                    break;
                case 50:
                    this.gzippedDownloadUrl = x0.readString();
                    this.hasGzippedDownloadUrl = true;
                    break;
                case 58:
                    if (this.patchData == null) {
                        this.patchData = new AndroidAppPatchData();
                    }
                    x0.readMessage(this.patchData);
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

    public static SplitDeliveryData[] emptyArray() {
        if (_emptyArray == null) {
            synchronized (InternalNano.LAZY_INIT_LOCK) {
                if (_emptyArray == null) {
                    _emptyArray = new SplitDeliveryData[0];
                }
            }
        }
        return _emptyArray;
    }

    public SplitDeliveryData() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasId || !this.id.equals("")) {
            output.writeString(1, this.id);
        }
        if (this.hasDownloadSize || this.downloadSize != 0) {
            output.writeInt64(2, this.downloadSize);
        }
        if (this.hasGzippedDownloadSize || this.gzippedDownloadSize != 0) {
            output.writeInt64(3, this.gzippedDownloadSize);
        }
        if (this.hasSignature || !this.signature.equals("")) {
            output.writeString(4, this.signature);
        }
        if (this.hasDownloadUrl || !this.downloadUrl.equals("")) {
            output.writeString(5, this.downloadUrl);
        }
        if (this.hasGzippedDownloadUrl || !this.gzippedDownloadUrl.equals("")) {
            output.writeString(6, this.gzippedDownloadUrl);
        }
        if (this.patchData != null) {
            output.writeMessage(7, this.patchData);
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasId || !this.id.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(1, this.id);
        }
        if (this.hasDownloadSize || this.downloadSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(2, this.downloadSize);
        }
        if (this.hasGzippedDownloadSize || this.gzippedDownloadSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(3, this.gzippedDownloadSize);
        }
        if (this.hasSignature || !this.signature.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(4, this.signature);
        }
        if (this.hasDownloadUrl || !this.downloadUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(5, this.downloadUrl);
        }
        if (this.hasGzippedDownloadUrl || !this.gzippedDownloadUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(6, this.gzippedDownloadUrl);
        }
        if (this.patchData != null) {
            return size + CodedOutputByteBufferNano.computeMessageSize(7, this.patchData);
        }
        return size;
    }
}
