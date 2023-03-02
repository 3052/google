package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class GetReviewsResponse extends MessageNano {
    public Review[] review = Review.emptyArray();
    public ReviewTip tip = null;
    public long matchingCount = 0;
    public boolean hasMatchingCount = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 10);
                    if (this.review == null) {
                        length = 0;
                    } else {
                        length = this.review.length;
                    }
                    Review[] reviewArr = new Review[repeatedFieldArrayLength + length];
                    if (length != 0) {
                        System.arraycopy(this.review, 0, reviewArr, 0, length);
                    }
                    while (length < reviewArr.length - 1) {
                        reviewArr[length] = new Review();
                        x0.readMessage(reviewArr[length]);
                        x0.readTag();
                        length++;
                    }
                    reviewArr[length] = new Review();
                    x0.readMessage(reviewArr[length]);
                    this.review = reviewArr;
                    break;
                case 16:
                    this.matchingCount = x0.readRawVarint64();
                    this.hasMatchingCount = true;
                    break;
                case 26:
                    if (this.tip == null) {
                        this.tip = new ReviewTip();
                    }
                    x0.readMessage(this.tip);
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

    public GetReviewsResponse() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.review != null && this.review.length > 0) {
            for (int i = 0; i < this.review.length; i++) {
                Review element = this.review[i];
                if (element != null) {
                    output.writeMessage(1, element);
                }
            }
        }
        if (this.hasMatchingCount || this.matchingCount != 0) {
            output.writeInt64(2, this.matchingCount);
        }
        if (this.tip != null) {
            output.writeMessage(3, this.tip);
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.review != null && this.review.length > 0) {
            for (int i = 0; i < this.review.length; i++) {
                Review element = this.review[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(1, element);
                }
            }
        }
        if (this.hasMatchingCount || this.matchingCount != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(2, this.matchingCount);
        }
        if (this.tip != null) {
            return size + CodedOutputByteBufferNano.computeMessageSize(3, this.tip);
        }
        return size;
    }
}
