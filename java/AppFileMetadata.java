package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class AppFileMetadata extends MessageNano {
    private static volatile AppFileMetadata[] _emptyArray;
    public int fileType = 0;
    public boolean hasFileType = false;
    public int versionCode = 0;
    public boolean hasVersionCode = false;
    public long size = 0;
    public boolean hasSize = false;
    public String signature = "";
    public boolean hasSignature = false;
    public long compressedSize = 0;
    public boolean hasCompressedSize = false;
    public String downloadUrl = "";
    public boolean hasDownloadUrl = false;
    public String compressedDownloadUrl = "";
    public boolean hasCompressedDownloadUrl = false;
    public AndroidAppPatchData patchData = null;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 8:
                    int readRawVarint32 = x0.readRawVarint32();
                    switch (readRawVarint32) {
                        case 0:
                        case 1:
                            this.fileType = readRawVarint32;
                            this.hasFileType = true;
                            continue;
                    }
                case 16:
                    this.versionCode = x0.readRawVarint32();
                    this.hasVersionCode = true;
                    break;
                case 24:
                    this.size = x0.readRawVarint64();
                    this.hasSize = true;
                    break;
                case 34:
                    this.downloadUrl = x0.readString();
                    this.hasDownloadUrl = true;
                    break;
                case 42:
                    if (this.patchData == null) {
                        this.patchData = new AndroidAppPatchData();
                    }
                    x0.readMessage(this.patchData);
                    break;
                case 48:
                    this.compressedSize = x0.readRawVarint64();
                    this.hasCompressedSize = true;
                    break;
                case 58:
                    this.compressedDownloadUrl = x0.readString();
                    this.hasCompressedDownloadUrl = true;
                    break;
                case 66:
                    this.signature = x0.readString();
                    this.hasSignature = true;
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

    public static AppFileMetadata[] emptyArray() {
        if (_emptyArray == null) {
            synchronized (InternalNano.LAZY_INIT_LOCK) {
                if (_emptyArray == null) {
                    _emptyArray = new AppFileMetadata[0];
                }
            }
        }
        return _emptyArray;
    }

    public AppFileMetadata() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.fileType != 0 || this.hasFileType) {
            output.writeInt32(1, this.fileType);
        }
        if (this.hasVersionCode || this.versionCode != 0) {
            output.writeInt32(2, this.versionCode);
        }
        if (this.hasSize || this.size != 0) {
            output.writeInt64(3, this.size);
        }
        if (this.hasDownloadUrl || !this.downloadUrl.equals("")) {
            output.writeString(4, this.downloadUrl);
        }
        if (this.patchData != null) {
            output.writeMessage(5, this.patchData);
        }
        if (this.hasCompressedSize || this.compressedSize != 0) {
            output.writeInt64(6, this.compressedSize);
        }
        if (this.hasCompressedDownloadUrl || !this.compressedDownloadUrl.equals("")) {
            output.writeString(7, this.compressedDownloadUrl);
        }
        if (this.hasSignature || !this.signature.equals("")) {
            output.writeString(8, this.signature);
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.fileType != 0 || this.hasFileType) {
            size += CodedOutputByteBufferNano.computeInt32Size(1, this.fileType);
        }
        if (this.hasVersionCode || this.versionCode != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(2, this.versionCode);
        }
        if (this.hasSize || this.size != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(3, this.size);
        }
        if (this.hasDownloadUrl || !this.downloadUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(4, this.downloadUrl);
        }
        if (this.patchData != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(5, this.patchData);
        }
        if (this.hasCompressedSize || this.compressedSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(6, this.compressedSize);
        }
        if (this.hasCompressedDownloadUrl || !this.compressedDownloadUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(7, this.compressedDownloadUrl);
        }
        if (this.hasSignature || !this.signature.equals("")) {
            return size + CodedOutputByteBufferNano.computeStringSize(8, this.signature);
        }
        return size;
    }
}
