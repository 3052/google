package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
import java.util.Arrays;
/* loaded from: classes.dex */
public final class PreFetch extends MessageNano {
    private static volatile PreFetch[] _emptyArray;
    public String url = "";
    public boolean hasUrl = false;
    public byte[] response = WireFormatNano.EMPTY_BYTES;
    public boolean hasResponse = false;
    public String etag = "";
    public boolean hasEtag = false;
    public long ttl = 0;
    public boolean hasTtl = false;
    public long softTtl = 0;
    public boolean hasSoftTtl = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    this.url = x0.readString();
                    this.hasUrl = true;
                    break;
                case 18:
                    this.response = x0.readBytes();
                    this.hasResponse = true;
                    break;
                case 26:
                    this.etag = x0.readString();
                    this.hasEtag = true;
                    break;
                case 32:
                    this.ttl = x0.readRawVarint64();
                    this.hasTtl = true;
                    break;
                case 40:
                    this.softTtl = x0.readRawVarint64();
                    this.hasSoftTtl = true;
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

    public static PreFetch[] emptyArray() {
        if (_emptyArray == null) {
            synchronized (InternalNano.LAZY_INIT_LOCK) {
                if (_emptyArray == null) {
                    _emptyArray = new PreFetch[0];
                }
            }
        }
        return _emptyArray;
    }

    public PreFetch() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasUrl || !this.url.equals("")) {
            output.writeString(1, this.url);
        }
        if (this.hasResponse || !Arrays.equals(this.response, WireFormatNano.EMPTY_BYTES)) {
            output.writeBytes(2, this.response);
        }
        if (this.hasEtag || !this.etag.equals("")) {
            output.writeString(3, this.etag);
        }
        if (this.hasTtl || this.ttl != 0) {
            output.writeInt64(4, this.ttl);
        }
        if (this.hasSoftTtl || this.softTtl != 0) {
            output.writeInt64(5, this.softTtl);
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasUrl || !this.url.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(1, this.url);
        }
        if (this.hasResponse || !Arrays.equals(this.response, WireFormatNano.EMPTY_BYTES)) {
            size += CodedOutputByteBufferNano.computeBytesSize(2, this.response);
        }
        if (this.hasEtag || !this.etag.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(3, this.etag);
        }
        if (this.hasTtl || this.ttl != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(4, this.ttl);
        }
        if (this.hasSoftTtl || this.softTtl != 0) {
            return size + CodedOutputByteBufferNano.computeInt64Size(5, this.softTtl);
        }
        return size;
    }
}
