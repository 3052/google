package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class FileMetadata extends MessageNano {
    private static volatile FileMetadata[] _emptyArray;
    public int fileType = 0;
    public boolean hasFileType = false;
    public int versionCode = 0;
    public boolean hasVersionCode = false;
    public long size = 0;
    public boolean hasSize = false;
    public long compressedSize = 0;
    public boolean hasCompressedSize = false;
    public PatchDetails[] patchDetails = PatchDetails.emptyArray();
    public String splitId = "";
    public boolean hasSplitId = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
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
                        case 2:
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
                    this.splitId = x0.readString();
                    this.hasSplitId = true;
                    break;
                case 40:
                    this.compressedSize = x0.readRawVarint64();
                    this.hasCompressedSize = true;
                    break;
                case 50:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 50);
                    if (this.patchDetails == null) {
                        length = 0;
                    } else {
                        length = this.patchDetails.length;
                    }
                    PatchDetails[] patchDetailsArr = new PatchDetails[repeatedFieldArrayLength + length];
                    if (length != 0) {
                        System.arraycopy(this.patchDetails, 0, patchDetailsArr, 0, length);
                    }
                    while (length < patchDetailsArr.length - 1) {
                        patchDetailsArr[length] = new PatchDetails();
                        x0.readMessage(patchDetailsArr[length]);
                        x0.readTag();
                        length++;
                    }
                    patchDetailsArr[length] = new PatchDetails();
                    x0.readMessage(patchDetailsArr[length]);
                    this.patchDetails = patchDetailsArr;
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

    public static FileMetadata[] emptyArray() {
        if (_emptyArray == null) {
            synchronized (InternalNano.LAZY_INIT_LOCK) {
                if (_emptyArray == null) {
                    _emptyArray = new FileMetadata[0];
                }
            }
        }
        return _emptyArray;
    }

    public FileMetadata() {
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
        if (this.hasSplitId || !this.splitId.equals("")) {
            output.writeString(4, this.splitId);
        }
        if (this.hasCompressedSize || this.compressedSize != 0) {
            output.writeInt64(5, this.compressedSize);
        }
        if (this.patchDetails != null && this.patchDetails.length > 0) {
            for (int i = 0; i < this.patchDetails.length; i++) {
                PatchDetails element = this.patchDetails[i];
                if (element != null) {
                    output.writeMessage(6, element);
                }
            }
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
        if (this.hasSplitId || !this.splitId.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(4, this.splitId);
        }
        if (this.hasCompressedSize || this.compressedSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(5, this.compressedSize);
        }
        if (this.patchDetails != null && this.patchDetails.length > 0) {
            for (int i = 0; i < this.patchDetails.length; i++) {
                PatchDetails element = this.patchDetails[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(6, element);
                }
            }
        }
        return size;
    }
}
