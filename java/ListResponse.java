package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class ListResponse extends MessageNano {
    public Bucket[] bucket = Bucket.emptyArray();
    public DocV2[] doc = DocV2.emptyArray();

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
        int length2;
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 10);
                    if (this.bucket == null) {
                        length2 = 0;
                    } else {
                        length2 = this.bucket.length;
                    }
                    Bucket[] bucketArr = new Bucket[repeatedFieldArrayLength + length2];
                    if (length2 != 0) {
                        System.arraycopy(this.bucket, 0, bucketArr, 0, length2);
                    }
                    while (length2 < bucketArr.length - 1) {
                        bucketArr[length2] = new Bucket();
                        x0.readMessage(bucketArr[length2]);
                        x0.readTag();
                        length2++;
                    }
                    bucketArr[length2] = new Bucket();
                    x0.readMessage(bucketArr[length2]);
                    this.bucket = bucketArr;
                    break;
                case 18:
                    int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 18);
                    if (this.doc == null) {
                        length = 0;
                    } else {
                        length = this.doc.length;
                    }
                    DocV2[] docV2Arr = new DocV2[repeatedFieldArrayLength2 + length];
                    if (length != 0) {
                        System.arraycopy(this.doc, 0, docV2Arr, 0, length);
                    }
                    while (length < docV2Arr.length - 1) {
                        docV2Arr[length] = new DocV2();
                        x0.readMessage(docV2Arr[length]);
                        x0.readTag();
                        length++;
                    }
                    docV2Arr[length] = new DocV2();
                    x0.readMessage(docV2Arr[length]);
                    this.doc = docV2Arr;
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

    public ListResponse() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.bucket != null && this.bucket.length > 0) {
            for (int i = 0; i < this.bucket.length; i++) {
                Bucket element = this.bucket[i];
                if (element != null) {
                    output.writeMessage(1, element);
                }
            }
        }
        if (this.doc != null && this.doc.length > 0) {
            for (int i2 = 0; i2 < this.doc.length; i2++) {
                DocV2 element2 = this.doc[i2];
                if (element2 != null) {
                    output.writeMessage(2, element2);
                }
            }
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.bucket != null && this.bucket.length > 0) {
            for (int i = 0; i < this.bucket.length; i++) {
                Bucket element = this.bucket[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(1, element);
                }
            }
        }
        if (this.doc != null && this.doc.length > 0) {
            for (int i2 = 0; i2 < this.doc.length; i2++) {
                DocV2 element2 = this.doc[i2];
                if (element2 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(2, element2);
                }
            }
        }
        return size;
    }
}
