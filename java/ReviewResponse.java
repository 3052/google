package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class ReviewResponse extends MessageNano {
    public GetReviewsResponse getResponse = null;
    public CriticReviewsResponse criticReviewsResponse = null;
    public String nextPageUrl = "";
    public boolean hasNextPageUrl = false;
    public Review updatedReview = null;
    public String suggestionsListUrl = "";
    public boolean hasSuggestionsListUrl = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    if (this.getResponse == null) {
                        this.getResponse = new GetReviewsResponse();
                    }
                    x0.readMessage(this.getResponse);
                    break;
                case 18:
                    this.nextPageUrl = x0.readString();
                    this.hasNextPageUrl = true;
                    break;
                case 26:
                    if (this.updatedReview == null) {
                        this.updatedReview = new Review();
                    }
                    x0.readMessage(this.updatedReview);
                    break;
                case 34:
                    this.suggestionsListUrl = x0.readString();
                    this.hasSuggestionsListUrl = true;
                    break;
                case 42:
                    if (this.criticReviewsResponse == null) {
                        this.criticReviewsResponse = new CriticReviewsResponse();
                    }
                    x0.readMessage(this.criticReviewsResponse);
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

    public ReviewResponse() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.getResponse != null) {
            output.writeMessage(1, this.getResponse);
        }
        if (this.hasNextPageUrl || !this.nextPageUrl.equals("")) {
            output.writeString(2, this.nextPageUrl);
        }
        if (this.updatedReview != null) {
            output.writeMessage(3, this.updatedReview);
        }
        if (this.hasSuggestionsListUrl || !this.suggestionsListUrl.equals("")) {
            output.writeString(4, this.suggestionsListUrl);
        }
        if (this.criticReviewsResponse != null) {
            output.writeMessage(5, this.criticReviewsResponse);
        }
        super.writeTo(output);
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.getResponse != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(1, this.getResponse);
        }
        if (this.hasNextPageUrl || !this.nextPageUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(2, this.nextPageUrl);
        }
        if (this.updatedReview != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(3, this.updatedReview);
        }
        if (this.hasSuggestionsListUrl || !this.suggestionsListUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(4, this.suggestionsListUrl);
        }
        if (this.criticReviewsResponse != null) {
            return size + CodedOutputByteBufferNano.computeMessageSize(5, this.criticReviewsResponse);
        }
        return size;
    }
}
