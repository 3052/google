package com.google.android.finsky.protos;

import com.google.android.finsky.protos.Common;
import com.google.android.finsky.protos.Containers;
import com.google.android.finsky.protos.DocDetails;
import com.google.android.finsky.protos.FilterRules;
import com.google.android.finsky.protos.Rating;
import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
import java.util.Arrays;
/* loaded from: classes.dex */
public final class DocV2 extends MessageNano {
    private static volatile DocV2[] _emptyArray;
    public String docid = "";
    public boolean hasDocid = false;
    public String backendDocid = "";
    public boolean hasBackendDocid = false;
    public int docType = 1;
    public boolean hasDocType = false;
    public int backendId = 0;
    public boolean hasBackendId = false;
    public String title = "";
    public boolean hasTitle = false;
    public String subtitle = "";
    public boolean hasSubtitle = false;
    public String creator = "";
    public boolean hasCreator = false;
    public String descriptionHtml = "";
    public boolean hasDescriptionHtml = false;
    public String translatedDescriptionHtml = "";
    public boolean hasTranslatedDescriptionHtml = false;
    public String promotionalDescription = "";
    public boolean hasPromotionalDescription = false;
    public Common.Offer[] offer = Common.Offer.emptyArray();
    public FilterRules.Availability availability = null;
    public Common.Image[] image = Common.Image.emptyArray();
    public DocV2[] child = emptyArray();
    public Containers.ContainerMetadata containerMetadata = null;
    public DocDetails.DocumentDetails details = null;
    public DocDetails.ProductDetails productDetails = null;
    public Rating.AggregateRating aggregateRating = null;
    public Annotations annotations = null;
    public String detailsUrl = "";
    public boolean hasDetailsUrl = false;
    public String shareUrl = "";
    public boolean hasShareUrl = false;
    public String reviewsUrl = "";
    public boolean hasReviewsUrl = false;
    public String snippetsUrl = "";
    public boolean hasSnippetsUrl = false;
    public String backendUrl = "";
    public boolean hasBackendUrl = false;
    public String purchaseDetailsUrl = "";
    public boolean hasPurchaseDetailsUrl = false;
    public boolean detailsReusable = false;
    public boolean hasDetailsReusable = false;
    public byte[] serverLogsCookie = WireFormatNano.EMPTY_BYTES;
    public boolean hasServerLogsCookie = false;
    public boolean mature = false;
    public boolean hasMature = false;
    public boolean availableForPreregistration = false;
    public boolean hasAvailableForPreregistration = false;
    public ReviewTip[] tip = ReviewTip.emptyArray();
    public boolean forceShareability = false;
    public boolean hasForceShareability = false;
    public boolean useWishlistAsPrimaryAction = false;
    public boolean hasUseWishlistAsPrimaryAction = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
        int length2;
        int length3;
        int length4;
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    this.docid = x0.readString();
                    this.hasDocid = true;
                    break;
                case 18:
                    this.backendDocid = x0.readString();
                    this.hasBackendDocid = true;
                    break;
                case 24:
                    int readRawVarint32 = x0.readRawVarint32();
                    switch (readRawVarint32) {
                        case 1:
                        case 2:
                        case 3:
                        case 4:
                        case 5:
                        case 6:
                        case 7:
                        case 8:
                        case 9:
                        case 10:
                        case 11:
                        case 12:
                        case 13:
                        case 14:
                        case 15:
                        case 16:
                        case 17:
                        case 18:
                        case 19:
                        case 20:
                        case 21:
                        case 22:
                        case 23:
                        case 24:
                        case 25:
                        case 26:
                        case 27:
                        case 28:
                        case 29:
                        case 30:
                        case 31:
                        case 32:
                        case 33:
                        case 34:
                        case 35:
                        case 36:
                        case 37:
                        case 38:
                        case 39:
                        case 40:
                        case 41:
                        case 42:
                        case 43:
                        case 44:
                        case 45:
                        case 46:
                        case 47:
                        case 48:
                        case 49:
                            this.docType = readRawVarint32;
                            this.hasDocType = true;
                            continue;
                    }
                case 32:
                    int readRawVarint322 = x0.readRawVarint32();
                    switch (readRawVarint322) {
                        case 0:
                        case 1:
                        case 2:
                        case 3:
                        case 4:
                        case 5:
                        case 6:
                        case 7:
                        case 9:
                        case 10:
                        case 11:
                        case 12:
                        case 13:
                            this.backendId = readRawVarint322;
                            this.hasBackendId = true;
                            continue;
                    }
                case 42:
                    this.title = x0.readString();
                    this.hasTitle = true;
                    break;
                case 50:
                    this.creator = x0.readString();
                    this.hasCreator = true;
                    break;
                case 58:
                    this.descriptionHtml = x0.readString();
                    this.hasDescriptionHtml = true;
                    break;
                case 66:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 66);
                    if (this.offer == null) {
                        length4 = 0;
                    } else {
                        length4 = this.offer.length;
                    }
                    Common.Offer[] offerArr = new Common.Offer[repeatedFieldArrayLength + length4];
                    if (length4 != 0) {
                        System.arraycopy(this.offer, 0, offerArr, 0, length4);
                    }
                    while (length4 < offerArr.length - 1) {
                        offerArr[length4] = new Common.Offer();
                        x0.readMessage(offerArr[length4]);
                        x0.readTag();
                        length4++;
                    }
                    offerArr[length4] = new Common.Offer();
                    x0.readMessage(offerArr[length4]);
                    this.offer = offerArr;
                    break;
                case 74:
                    if (this.availability == null) {
                        this.availability = new FilterRules.Availability();
                    }
                    x0.readMessage(this.availability);
                    break;
                case 82:
                    int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 82);
                    if (this.image == null) {
                        length3 = 0;
                    } else {
                        length3 = this.image.length;
                    }
                    Common.Image[] imageArr = new Common.Image[repeatedFieldArrayLength2 + length3];
                    if (length3 != 0) {
                        System.arraycopy(this.image, 0, imageArr, 0, length3);
                    }
                    while (length3 < imageArr.length - 1) {
                        imageArr[length3] = new Common.Image();
                        x0.readMessage(imageArr[length3]);
                        x0.readTag();
                        length3++;
                    }
                    imageArr[length3] = new Common.Image();
                    x0.readMessage(imageArr[length3]);
                    this.image = imageArr;
                    break;
                case 90:
                    int repeatedFieldArrayLength3 = WireFormatNano.getRepeatedFieldArrayLength(x0, 90);
                    if (this.child == null) {
                        length2 = 0;
                    } else {
                        length2 = this.child.length;
                    }
                    DocV2[] docV2Arr = new DocV2[repeatedFieldArrayLength3 + length2];
                    if (length2 != 0) {
                        System.arraycopy(this.child, 0, docV2Arr, 0, length2);
                    }
                    while (length2 < docV2Arr.length - 1) {
                        docV2Arr[length2] = new DocV2();
                        x0.readMessage(docV2Arr[length2]);
                        x0.readTag();
                        length2++;
                    }
                    docV2Arr[length2] = new DocV2();
                    x0.readMessage(docV2Arr[length2]);
                    this.child = docV2Arr;
                    break;
                case 98:
                    if (this.containerMetadata == null) {
                        this.containerMetadata = new Containers.ContainerMetadata();
                    }
                    x0.readMessage(this.containerMetadata);
                    break;
                case 106:
                    if (this.details == null) {
                        this.details = new DocDetails.DocumentDetails();
                    }
                    x0.readMessage(this.details);
                    break;
                case 114:
                    if (this.aggregateRating == null) {
                        this.aggregateRating = new Rating.AggregateRating();
                    }
                    x0.readMessage(this.aggregateRating);
                    break;
                case 122:
                    if (this.annotations == null) {
                        this.annotations = new Annotations();
                    }
                    x0.readMessage(this.annotations);
                    break;
                case 130:
                    this.detailsUrl = x0.readString();
                    this.hasDetailsUrl = true;
                    break;
                case 138:
                    this.shareUrl = x0.readString();
                    this.hasShareUrl = true;
                    break;
                case 146:
                    this.reviewsUrl = x0.readString();
                    this.hasReviewsUrl = true;
                    break;
                case 154:
                    this.backendUrl = x0.readString();
                    this.hasBackendUrl = true;
                    break;
                case 162:
                    this.purchaseDetailsUrl = x0.readString();
                    this.hasPurchaseDetailsUrl = true;
                    break;
                case 168:
                    this.detailsReusable = x0.readBool();
                    this.hasDetailsReusable = true;
                    break;
                case 178:
                    this.subtitle = x0.readString();
                    this.hasSubtitle = true;
                    break;
                case 186:
                    this.translatedDescriptionHtml = x0.readString();
                    this.hasTranslatedDescriptionHtml = true;
                    break;
                case 194:
                    this.serverLogsCookie = x0.readBytes();
                    this.hasServerLogsCookie = true;
                    break;
                case 202:
                    if (this.productDetails == null) {
                        this.productDetails = new DocDetails.ProductDetails();
                    }
                    x0.readMessage(this.productDetails);
                    break;
                case 208:
                    this.mature = x0.readBool();
                    this.hasMature = true;
                    break;
                case 218:
                    this.promotionalDescription = x0.readString();
                    this.hasPromotionalDescription = true;
                    break;
                case 232:
                    this.availableForPreregistration = x0.readBool();
                    this.hasAvailableForPreregistration = true;
                    break;
                case 242:
                    int repeatedFieldArrayLength4 = WireFormatNano.getRepeatedFieldArrayLength(x0, 242);
                    if (this.tip == null) {
                        length = 0;
                    } else {
                        length = this.tip.length;
                    }
                    ReviewTip[] reviewTipArr = new ReviewTip[repeatedFieldArrayLength4 + length];
                    if (length != 0) {
                        System.arraycopy(this.tip, 0, reviewTipArr, 0, length);
                    }
                    while (length < reviewTipArr.length - 1) {
                        reviewTipArr[length] = new ReviewTip();
                        x0.readMessage(reviewTipArr[length]);
                        x0.readTag();
                        length++;
                    }
                    reviewTipArr[length] = new ReviewTip();
                    x0.readMessage(reviewTipArr[length]);
                    this.tip = reviewTipArr;
                    break;
                case 250:
                    this.snippetsUrl = x0.readString();
                    this.hasSnippetsUrl = true;
                    break;
                case 256:
                    this.forceShareability = x0.readBool();
                    this.hasForceShareability = true;
                    break;
                case 264:
                    this.useWishlistAsPrimaryAction = x0.readBool();
                    this.hasUseWishlistAsPrimaryAction = true;
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

    public static DocV2[] emptyArray() {
        if (_emptyArray == null) {
            synchronized (InternalNano.LAZY_INIT_LOCK) {
                if (_emptyArray == null) {
                    _emptyArray = new DocV2[0];
                }
            }
        }
        return _emptyArray;
    }

    public DocV2() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasDocid || !this.docid.equals("")) {
            output.writeString(1, this.docid);
        }
        if (this.hasBackendDocid || !this.backendDocid.equals("")) {
            output.writeString(2, this.backendDocid);
        }
        if (this.docType != 1 || this.hasDocType) {
            output.writeInt32(3, this.docType);
        }
        if (this.backendId != 0 || this.hasBackendId) {
            output.writeInt32(4, this.backendId);
        }
        if (this.hasTitle || !this.title.equals("")) {
            output.writeString(5, this.title);
        }
        if (this.hasCreator || !this.creator.equals("")) {
            output.writeString(6, this.creator);
        }
        if (this.hasDescriptionHtml || !this.descriptionHtml.equals("")) {
            output.writeString(7, this.descriptionHtml);
        }
        if (this.offer != null && this.offer.length > 0) {
            for (int i = 0; i < this.offer.length; i++) {
                Common.Offer element = this.offer[i];
                if (element != null) {
                    output.writeMessage(8, element);
                }
            }
        }
        if (this.availability != null) {
            output.writeMessage(9, this.availability);
        }
        if (this.image != null && this.image.length > 0) {
            for (int i2 = 0; i2 < this.image.length; i2++) {
                Common.Image element2 = this.image[i2];
                if (element2 != null) {
                    output.writeMessage(10, element2);
                }
            }
        }
        if (this.child != null && this.child.length > 0) {
            for (int i3 = 0; i3 < this.child.length; i3++) {
                DocV2 element3 = this.child[i3];
                if (element3 != null) {
                    output.writeMessage(11, element3);
                }
            }
        }
        if (this.containerMetadata != null) {
            output.writeMessage(12, this.containerMetadata);
        }
        if (this.details != null) {
            output.writeMessage(13, this.details);
        }
        if (this.aggregateRating != null) {
            output.writeMessage(14, this.aggregateRating);
        }
        if (this.annotations != null) {
            output.writeMessage(15, this.annotations);
        }
        if (this.hasDetailsUrl || !this.detailsUrl.equals("")) {
            output.writeString(16, this.detailsUrl);
        }
        if (this.hasShareUrl || !this.shareUrl.equals("")) {
            output.writeString(17, this.shareUrl);
        }
        if (this.hasReviewsUrl || !this.reviewsUrl.equals("")) {
            output.writeString(18, this.reviewsUrl);
        }
        if (this.hasBackendUrl || !this.backendUrl.equals("")) {
            output.writeString(19, this.backendUrl);
        }
        if (this.hasPurchaseDetailsUrl || !this.purchaseDetailsUrl.equals("")) {
            output.writeString(20, this.purchaseDetailsUrl);
        }
        if (this.hasDetailsReusable || this.detailsReusable) {
            output.writeBool(21, this.detailsReusable);
        }
        if (this.hasSubtitle || !this.subtitle.equals("")) {
            output.writeString(22, this.subtitle);
        }
        if (this.hasTranslatedDescriptionHtml || !this.translatedDescriptionHtml.equals("")) {
            output.writeString(23, this.translatedDescriptionHtml);
        }
        if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
            output.writeBytes(24, this.serverLogsCookie);
        }
        if (this.productDetails != null) {
            output.writeMessage(25, this.productDetails);
        }
        if (this.hasMature || this.mature) {
            output.writeBool(26, this.mature);
        }
        if (this.hasPromotionalDescription || !this.promotionalDescription.equals("")) {
            output.writeString(27, this.promotionalDescription);
        }
        if (this.hasAvailableForPreregistration || this.availableForPreregistration) {
            output.writeBool(29, this.availableForPreregistration);
        }
        if (this.tip != null && this.tip.length > 0) {
            for (int i4 = 0; i4 < this.tip.length; i4++) {
                ReviewTip element4 = this.tip[i4];
                if (element4 != null) {
                    output.writeMessage(30, element4);
                }
            }
        }
        if (this.hasSnippetsUrl || !this.snippetsUrl.equals("")) {
            output.writeString(31, this.snippetsUrl);
        }
        if (this.hasForceShareability || this.forceShareability) {
            output.writeBool(32, this.forceShareability);
        }
        if (this.hasUseWishlistAsPrimaryAction || this.useWishlistAsPrimaryAction) {
            output.writeBool(33, this.useWishlistAsPrimaryAction);
        }
        super.writeTo(output);
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasDocid || !this.docid.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(1, this.docid);
        }
        if (this.hasBackendDocid || !this.backendDocid.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(2, this.backendDocid);
        }
        if (this.docType != 1 || this.hasDocType) {
            size += CodedOutputByteBufferNano.computeInt32Size(3, this.docType);
        }
        if (this.backendId != 0 || this.hasBackendId) {
            size += CodedOutputByteBufferNano.computeInt32Size(4, this.backendId);
        }
        if (this.hasTitle || !this.title.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(5, this.title);
        }
        if (this.hasCreator || !this.creator.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(6, this.creator);
        }
        if (this.hasDescriptionHtml || !this.descriptionHtml.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(7, this.descriptionHtml);
        }
        if (this.offer != null && this.offer.length > 0) {
            for (int i = 0; i < this.offer.length; i++) {
                Common.Offer element = this.offer[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(8, element);
                }
            }
        }
        if (this.availability != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(9, this.availability);
        }
        if (this.image != null && this.image.length > 0) {
            for (int i2 = 0; i2 < this.image.length; i2++) {
                Common.Image element2 = this.image[i2];
                if (element2 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(10, element2);
                }
            }
        }
        if (this.child != null && this.child.length > 0) {
            for (int i3 = 0; i3 < this.child.length; i3++) {
                DocV2 element3 = this.child[i3];
                if (element3 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(11, element3);
                }
            }
        }
        if (this.containerMetadata != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(12, this.containerMetadata);
        }
        if (this.details != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(13, this.details);
        }
        if (this.aggregateRating != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(14, this.aggregateRating);
        }
        if (this.annotations != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(15, this.annotations);
        }
        if (this.hasDetailsUrl || !this.detailsUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(16, this.detailsUrl);
        }
        if (this.hasShareUrl || !this.shareUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(17, this.shareUrl);
        }
        if (this.hasReviewsUrl || !this.reviewsUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(18, this.reviewsUrl);
        }
        if (this.hasBackendUrl || !this.backendUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(19, this.backendUrl);
        }
        if (this.hasPurchaseDetailsUrl || !this.purchaseDetailsUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(20, this.purchaseDetailsUrl);
        }
        if (this.hasDetailsReusable || this.detailsReusable) {
            size += CodedOutputByteBufferNano.computeTagSize(21) + 1;
        }
        if (this.hasSubtitle || !this.subtitle.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(22, this.subtitle);
        }
        if (this.hasTranslatedDescriptionHtml || !this.translatedDescriptionHtml.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(23, this.translatedDescriptionHtml);
        }
        if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
            size += CodedOutputByteBufferNano.computeBytesSize(24, this.serverLogsCookie);
        }
        if (this.productDetails != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(25, this.productDetails);
        }
        if (this.hasMature || this.mature) {
            size += CodedOutputByteBufferNano.computeTagSize(26) + 1;
        }
        if (this.hasPromotionalDescription || !this.promotionalDescription.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(27, this.promotionalDescription);
        }
        if (this.hasAvailableForPreregistration || this.availableForPreregistration) {
            size += CodedOutputByteBufferNano.computeTagSize(29) + 1;
        }
        if (this.tip != null && this.tip.length > 0) {
            for (int i4 = 0; i4 < this.tip.length; i4++) {
                ReviewTip element4 = this.tip[i4];
                if (element4 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(30, element4);
                }
            }
        }
        if (this.hasSnippetsUrl || !this.snippetsUrl.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(31, this.snippetsUrl);
        }
        if (this.hasForceShareability || this.forceShareability) {
            size += CodedOutputByteBufferNano.computeTagSize(32) + 1;
        }
        if (this.hasUseWishlistAsPrimaryAction || this.useWishlistAsPrimaryAction) {
            return size + CodedOutputByteBufferNano.computeTagSize(33) + 1;
        }
        return size;
    }
}
