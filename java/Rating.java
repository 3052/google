package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public interface Rating {

    /* loaded from: classes.dex */
    public static final class AggregateRating extends MessageNano {
        public int type = 1;
        public boolean hasType = false;
        public float starRating = 0.0f;
        public boolean hasStarRating = false;
        public long ratingsCount = 0;
        public boolean hasRatingsCount = false;
        public long commentCount = 0;
        public boolean hasCommentCount = false;
        public long oneStarRatings = 0;
        public boolean hasOneStarRatings = false;
        public long twoStarRatings = 0;
        public boolean hasTwoStarRatings = false;
        public long threeStarRatings = 0;
        public boolean hasThreeStarRatings = false;
        public long fourStarRatings = 0;
        public boolean hasFourStarRatings = false;
        public long fiveStarRatings = 0;
        public boolean hasFiveStarRatings = false;
        public double bayesianMeanRating = 0.0d;
        public boolean hasBayesianMeanRating = false;
        public long thumbsUpCount = 0;
        public boolean hasThumbsUpCount = false;
        public long thumbsDownCount = 0;
        public boolean hasThumbsDownCount = false;
        public Tip[] tip = Tip.emptyArray();

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
                            case 1:
                            case 2:
                            case 3:
                                this.type = readRawVarint32;
                                this.hasType = true;
                                continue;
                        }
                    case 21:
                        this.starRating = Float.intBitsToFloat(x0.readRawLittleEndian32());
                        this.hasStarRating = true;
                        break;
                    case 24:
                        this.ratingsCount = x0.readRawVarint64();
                        this.hasRatingsCount = true;
                        break;
                    case 32:
                        this.oneStarRatings = x0.readRawVarint64();
                        this.hasOneStarRatings = true;
                        break;
                    case 40:
                        this.twoStarRatings = x0.readRawVarint64();
                        this.hasTwoStarRatings = true;
                        break;
                    case 48:
                        this.threeStarRatings = x0.readRawVarint64();
                        this.hasThreeStarRatings = true;
                        break;
                    case 56:
                        this.fourStarRatings = x0.readRawVarint64();
                        this.hasFourStarRatings = true;
                        break;
                    case 64:
                        this.fiveStarRatings = x0.readRawVarint64();
                        this.hasFiveStarRatings = true;
                        break;
                    case 72:
                        this.thumbsUpCount = x0.readRawVarint64();
                        this.hasThumbsUpCount = true;
                        break;
                    case 80:
                        this.thumbsDownCount = x0.readRawVarint64();
                        this.hasThumbsDownCount = true;
                        break;
                    case 88:
                        this.commentCount = x0.readRawVarint64();
                        this.hasCommentCount = true;
                        break;
                    case 97:
                        this.bayesianMeanRating = Double.longBitsToDouble(x0.readRawLittleEndian64());
                        this.hasBayesianMeanRating = true;
                        break;
                    case 106:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 106);
                        if (this.tip == null) {
                            length = 0;
                        } else {
                            length = this.tip.length;
                        }
                        Tip[] tipArr = new Tip[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.tip, 0, tipArr, 0, length);
                        }
                        while (length < tipArr.length - 1) {
                            tipArr[length] = new Tip();
                            x0.readMessage(tipArr[length]);
                            x0.readTag();
                            length++;
                        }
                        tipArr[length] = new Tip();
                        x0.readMessage(tipArr[length]);
                        this.tip = tipArr;
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

        public AggregateRating() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.type != 1 || this.hasType) {
                output.writeInt32(1, this.type);
            }
            if (this.hasStarRating || Float.floatToIntBits(this.starRating) != Float.floatToIntBits(0.0f)) {
                output.writeFloat(2, this.starRating);
            }
            if (this.hasRatingsCount || this.ratingsCount != 0) {
                output.writeUInt64(3, this.ratingsCount);
            }
            if (this.hasOneStarRatings || this.oneStarRatings != 0) {
                output.writeUInt64(4, this.oneStarRatings);
            }
            if (this.hasTwoStarRatings || this.twoStarRatings != 0) {
                output.writeUInt64(5, this.twoStarRatings);
            }
            if (this.hasThreeStarRatings || this.threeStarRatings != 0) {
                output.writeUInt64(6, this.threeStarRatings);
            }
            if (this.hasFourStarRatings || this.fourStarRatings != 0) {
                output.writeUInt64(7, this.fourStarRatings);
            }
            if (this.hasFiveStarRatings || this.fiveStarRatings != 0) {
                output.writeUInt64(8, this.fiveStarRatings);
            }
            if (this.hasThumbsUpCount || this.thumbsUpCount != 0) {
                output.writeUInt64(9, this.thumbsUpCount);
            }
            if (this.hasThumbsDownCount || this.thumbsDownCount != 0) {
                output.writeUInt64(10, this.thumbsDownCount);
            }
            if (this.hasCommentCount || this.commentCount != 0) {
                output.writeUInt64(11, this.commentCount);
            }
            if (this.hasBayesianMeanRating || Double.doubleToLongBits(this.bayesianMeanRating) != Double.doubleToLongBits(0.0d)) {
                output.writeDouble(12, this.bayesianMeanRating);
            }
            if (this.tip != null && this.tip.length > 0) {
                for (int i = 0; i < this.tip.length; i++) {
                    Tip element = this.tip[i];
                    if (element != null) {
                        output.writeMessage(13, element);
                    }
                }
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.type != 1 || this.hasType) {
                size += CodedOutputByteBufferNano.computeInt32Size(1, this.type);
            }
            if (this.hasStarRating || Float.floatToIntBits(this.starRating) != Float.floatToIntBits(0.0f)) {
                size += CodedOutputByteBufferNano.computeTagSize(2) + 4;
            }
            if (this.hasRatingsCount || this.ratingsCount != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(3, this.ratingsCount);
            }
            if (this.hasOneStarRatings || this.oneStarRatings != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(4, this.oneStarRatings);
            }
            if (this.hasTwoStarRatings || this.twoStarRatings != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(5, this.twoStarRatings);
            }
            if (this.hasThreeStarRatings || this.threeStarRatings != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(6, this.threeStarRatings);
            }
            if (this.hasFourStarRatings || this.fourStarRatings != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(7, this.fourStarRatings);
            }
            if (this.hasFiveStarRatings || this.fiveStarRatings != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(8, this.fiveStarRatings);
            }
            if (this.hasThumbsUpCount || this.thumbsUpCount != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(9, this.thumbsUpCount);
            }
            if (this.hasThumbsDownCount || this.thumbsDownCount != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(10, this.thumbsDownCount);
            }
            if (this.hasCommentCount || this.commentCount != 0) {
                size += CodedOutputByteBufferNano.computeUInt64Size(11, this.commentCount);
            }
            if (this.hasBayesianMeanRating || Double.doubleToLongBits(this.bayesianMeanRating) != Double.doubleToLongBits(0.0d)) {
                size += CodedOutputByteBufferNano.computeTagSize(12) + 8;
            }
            if (this.tip != null && this.tip.length > 0) {
                for (int i = 0; i < this.tip.length; i++) {
                    Tip element = this.tip[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(13, element);
                    }
                }
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class Tip extends MessageNano {
        private static volatile Tip[] _emptyArray;
        public String tipId = "";
        public boolean hasTipId = false;
        public String text = "";
        public boolean hasText = false;
        public int polarity = 0;
        public boolean hasPolarity = false;
        public long reviewCount = 0;
        public boolean hasReviewCount = false;
        public String language = "";
        public boolean hasLanguage = false;
        public String[] snippetReviewId = WireFormatNano.EMPTY_STRING_ARRAY;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.tipId = x0.readString();
                        this.hasTipId = true;
                        break;
                    case 18:
                        this.text = x0.readString();
                        this.hasText = true;
                        break;
                    case 24:
                        int readRawVarint32 = x0.readRawVarint32();
                        switch (readRawVarint32) {
                            case 0:
                            case 1:
                            case 2:
                                this.polarity = readRawVarint32;
                                this.hasPolarity = true;
                                continue;
                        }
                    case 32:
                        this.reviewCount = x0.readRawVarint64();
                        this.hasReviewCount = true;
                        break;
                    case 42:
                        this.language = x0.readString();
                        this.hasLanguage = true;
                        break;
                    case 50:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 50);
                        int length = this.snippetReviewId == null ? 0 : this.snippetReviewId.length;
                        String[] strArr = new String[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.snippetReviewId, 0, strArr, 0, length);
                        }
                        while (length < strArr.length - 1) {
                            strArr[length] = x0.readString();
                            x0.readTag();
                            length++;
                        }
                        strArr[length] = x0.readString();
                        this.snippetReviewId = strArr;
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

        public static Tip[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new Tip[0];
                    }
                }
            }
            return _emptyArray;
        }

        public Tip() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasTipId || !this.tipId.equals("")) {
                output.writeString(1, this.tipId);
            }
            if (this.hasText || !this.text.equals("")) {
                output.writeString(2, this.text);
            }
            if (this.polarity != 0 || this.hasPolarity) {
                output.writeInt32(3, this.polarity);
            }
            if (this.hasReviewCount || this.reviewCount != 0) {
                output.writeInt64(4, this.reviewCount);
            }
            if (this.hasLanguage || !this.language.equals("")) {
                output.writeString(5, this.language);
            }
            if (this.snippetReviewId != null && this.snippetReviewId.length > 0) {
                for (int i = 0; i < this.snippetReviewId.length; i++) {
                    String element = this.snippetReviewId[i];
                    if (element != null) {
                        output.writeString(6, element);
                    }
                }
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasTipId || !this.tipId.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.tipId);
            }
            if (this.hasText || !this.text.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.text);
            }
            if (this.polarity != 0 || this.hasPolarity) {
                size += CodedOutputByteBufferNano.computeInt32Size(3, this.polarity);
            }
            if (this.hasReviewCount || this.reviewCount != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(4, this.reviewCount);
            }
            if (this.hasLanguage || !this.language.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(5, this.language);
            }
            if (this.snippetReviewId == null || this.snippetReviewId.length <= 0) {
                return size;
            }
            int dataCount = 0;
            int dataSize = 0;
            for (int i = 0; i < this.snippetReviewId.length; i++) {
                String element = this.snippetReviewId[i];
                if (element != null) {
                    dataCount++;
                    dataSize += CodedOutputByteBufferNano.computeStringSizeNoTag(element);
                }
            }
            return size + dataSize + (dataCount * 1);
        }
    }
}
