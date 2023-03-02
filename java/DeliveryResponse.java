package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class DeliveryResponse extends MessageNano {
    public int status = 1;
    public boolean hasStatus = false;
    public AndroidAppDeliveryData appDeliveryData = null;

    public DeliveryResponse() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.status != 1 || this.hasStatus) {
            output.writeInt32(1, this.status);
        }
        if (this.appDeliveryData != null) {
            output.writeMessage(2, this.appDeliveryData);
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.status != 1 || this.hasStatus) {
            size += CodedOutputByteBufferNano.computeInt32Size(1, this.status);
        }
        if (this.appDeliveryData != null) {
            return size + CodedOutputByteBufferNano.computeMessageSize(2, this.appDeliveryData);
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
                case 8:
                    int readRawVarint32 = x0.readRawVarint32();
                    switch (readRawVarint32) {
                        case 1:
                        case 2:
                        case 3:
                        case 4:
                        case 5:
                        case 6:
                            this.status = readRawVarint32;
                            this.hasStatus = true;
                            continue;
                    }
                case 18:
                    if (this.appDeliveryData == null) {
                        this.appDeliveryData = new AndroidAppDeliveryData();
                    }
                    x0.readMessage(this.appDeliveryData);
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
}
