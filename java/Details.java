package com.google.android.finsky.protos;

import com.google.android.finsky.protos.Common;
import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
import java.util.Arrays;
/* loaded from: classes.dex */
public interface Details {

    /* loaded from: classes.dex */
    public static final class DetailsResponse extends MessageNano {
        public DocV1 docV1 = null;
        public DocV2 docV2 = null;
        public Review userReview = null;
        public String footerHtml = "";
        public boolean hasFooterHtml = false;
        public byte[] serverLogsCookie = WireFormatNano.EMPTY_BYTES;
        public boolean hasServerLogsCookie = false;
        public DiscoveryBadge[] discoveryBadge = DiscoveryBadge.emptyArray();
        public boolean enableReviews = true;
        public boolean hasEnableReviews = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        if (this.docV1 == null) {
                            this.docV1 = new DocV1();
                        }
                        x0.readMessage(this.docV1);
                        break;
                    case 26:
                        if (this.userReview == null) {
                            this.userReview = new Review();
                        }
                        x0.readMessage(this.userReview);
                        break;
                    case 34:
                        if (this.docV2 == null) {
                            this.docV2 = new DocV2();
                        }
                        x0.readMessage(this.docV2);
                        break;
                    case 42:
                        this.footerHtml = x0.readString();
                        this.hasFooterHtml = true;
                        break;
                    case 50:
                        this.serverLogsCookie = x0.readBytes();
                        this.hasServerLogsCookie = true;
                        break;
                    case 58:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 58);
                        if (this.discoveryBadge == null) {
                            length = 0;
                        } else {
                            length = this.discoveryBadge.length;
                        }
                        DiscoveryBadge[] discoveryBadgeArr = new DiscoveryBadge[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.discoveryBadge, 0, discoveryBadgeArr, 0, length);
                        }
                        while (length < discoveryBadgeArr.length - 1) {
                            discoveryBadgeArr[length] = new DiscoveryBadge();
                            x0.readMessage(discoveryBadgeArr[length]);
                            x0.readTag();
                            length++;
                        }
                        discoveryBadgeArr[length] = new DiscoveryBadge();
                        x0.readMessage(discoveryBadgeArr[length]);
                        this.discoveryBadge = discoveryBadgeArr;
                        break;
                    case 64:
                        this.enableReviews = x0.readBool();
                        this.hasEnableReviews = true;
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

        public DetailsResponse() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.docV1 != null) {
                output.writeMessage(1, this.docV1);
            }
            if (this.userReview != null) {
                output.writeMessage(3, this.userReview);
            }
            if (this.docV2 != null) {
                output.writeMessage(4, this.docV2);
            }
            if (this.hasFooterHtml || !this.footerHtml.equals("")) {
                output.writeString(5, this.footerHtml);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                output.writeBytes(6, this.serverLogsCookie);
            }
            if (this.discoveryBadge != null && this.discoveryBadge.length > 0) {
                for (int i = 0; i < this.discoveryBadge.length; i++) {
                    DiscoveryBadge element = this.discoveryBadge[i];
                    if (element != null) {
                        output.writeMessage(7, element);
                    }
                }
            }
            if (this.hasEnableReviews || !this.enableReviews) {
                output.writeBool(8, this.enableReviews);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.docV1 != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.docV1);
            }
            if (this.userReview != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.userReview);
            }
            if (this.docV2 != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(4, this.docV2);
            }
            if (this.hasFooterHtml || !this.footerHtml.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(5, this.footerHtml);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                size += CodedOutputByteBufferNano.computeBytesSize(6, this.serverLogsCookie);
            }
            if (this.discoveryBadge != null && this.discoveryBadge.length > 0) {
                for (int i = 0; i < this.discoveryBadge.length; i++) {
                    DiscoveryBadge element = this.discoveryBadge[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(7, element);
                    }
                }
            }
            if (this.hasEnableReviews || !this.enableReviews) {
                return size + CodedOutputByteBufferNano.computeTagSize(8) + 1;
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class BulkDetailsRequest extends MessageNano {
        public String[] docid = WireFormatNano.EMPTY_STRING_ARRAY;
        public int[] installedVersionCode = WireFormatNano.EMPTY_INT_ARRAY;
        public boolean includeChildDocs = true;
        public boolean hasIncludeChildDocs = false;
        public boolean includeDetails = false;
        public boolean hasIncludeDetails = false;
        public String sourcePackageName = "";
        public boolean hasSourcePackageName = false;

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
                        int length2 = this.docid == null ? 0 : this.docid.length;
                        String[] strArr = new String[repeatedFieldArrayLength + length2];
                        if (length2 != 0) {
                            System.arraycopy(this.docid, 0, strArr, 0, length2);
                        }
                        while (length2 < strArr.length - 1) {
                            strArr[length2] = x0.readString();
                            x0.readTag();
                            length2++;
                        }
                        strArr[length2] = x0.readString();
                        this.docid = strArr;
                        break;
                    case 16:
                        this.includeChildDocs = x0.readBool();
                        this.hasIncludeChildDocs = true;
                        break;
                    case 24:
                        this.includeDetails = x0.readBool();
                        this.hasIncludeDetails = true;
                        break;
                    case 34:
                        this.sourcePackageName = x0.readString();
                        this.hasSourcePackageName = true;
                        break;
                    case 56:
                        int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 56);
                        if (this.installedVersionCode == null) {
                            length = 0;
                        } else {
                            length = this.installedVersionCode.length;
                        }
                        int[] iArr = new int[repeatedFieldArrayLength2 + length];
                        if (length != 0) {
                            System.arraycopy(this.installedVersionCode, 0, iArr, 0, length);
                        }
                        while (length < iArr.length - 1) {
                            iArr[length] = x0.readRawVarint32();
                            x0.readTag();
                            length++;
                        }
                        iArr[length] = x0.readRawVarint32();
                        this.installedVersionCode = iArr;
                        break;
                    case 58:
                        int pushLimit = x0.pushLimit(x0.readRawVarint32());
                        int position = x0.getPosition();
                        int i = 0;
                        while (x0.getBytesUntilLimit() > 0) {
                            x0.readRawVarint32();
                            i++;
                        }
                        x0.rewindToPosition(position);
                        int length3 = this.installedVersionCode == null ? 0 : this.installedVersionCode.length;
                        int[] iArr2 = new int[i + length3];
                        if (length3 != 0) {
                            System.arraycopy(this.installedVersionCode, 0, iArr2, 0, length3);
                        }
                        while (length3 < iArr2.length) {
                            iArr2[length3] = x0.readRawVarint32();
                            length3++;
                        }
                        this.installedVersionCode = iArr2;
                        x0.popLimit(pushLimit);
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

        public BulkDetailsRequest() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.docid != null && this.docid.length > 0) {
                for (int i = 0; i < this.docid.length; i++) {
                    String element = this.docid[i];
                    if (element != null) {
                        output.writeString(1, element);
                    }
                }
            }
            if (this.hasIncludeChildDocs || !this.includeChildDocs) {
                output.writeBool(2, this.includeChildDocs);
            }
            if (this.hasIncludeDetails || this.includeDetails) {
                output.writeBool(3, this.includeDetails);
            }
            if (this.hasSourcePackageName || !this.sourcePackageName.equals("")) {
                output.writeString(4, this.sourcePackageName);
            }
            if (this.installedVersionCode != null && this.installedVersionCode.length > 0) {
                for (int i2 = 0; i2 < this.installedVersionCode.length; i2++) {
                    output.writeInt32(7, this.installedVersionCode[i2]);
                }
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.docid != null && this.docid.length > 0) {
                int dataCount = 0;
                int dataSize = 0;
                for (int i = 0; i < this.docid.length; i++) {
                    String element = this.docid[i];
                    if (element != null) {
                        dataCount++;
                        dataSize += CodedOutputByteBufferNano.computeStringSizeNoTag(element);
                    }
                }
                size = size + dataSize + (dataCount * 1);
            }
            if (this.hasIncludeChildDocs || !this.includeChildDocs) {
                size += CodedOutputByteBufferNano.computeTagSize(2) + 1;
            }
            if (this.hasIncludeDetails || this.includeDetails) {
                size += CodedOutputByteBufferNano.computeTagSize(3) + 1;
            }
            if (this.hasSourcePackageName || !this.sourcePackageName.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(4, this.sourcePackageName);
            }
            if (this.installedVersionCode == null || this.installedVersionCode.length <= 0) {
                return size;
            }
            int dataSize2 = 0;
            for (int i2 = 0; i2 < this.installedVersionCode.length; i2++) {
                dataSize2 += CodedOutputByteBufferNano.computeInt32SizeNoTag(this.installedVersionCode[i2]);
            }
            return size + dataSize2 + (this.installedVersionCode.length * 1);
        }
    }

    /* loaded from: classes.dex */
    public static final class BulkDetailsResponse extends MessageNano {
        public BulkDetailsEntry[] entry = BulkDetailsEntry.emptyArray();

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
                        if (this.entry == null) {
                            length = 0;
                        } else {
                            length = this.entry.length;
                        }
                        BulkDetailsEntry[] bulkDetailsEntryArr = new BulkDetailsEntry[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.entry, 0, bulkDetailsEntryArr, 0, length);
                        }
                        while (length < bulkDetailsEntryArr.length - 1) {
                            bulkDetailsEntryArr[length] = new BulkDetailsEntry();
                            x0.readMessage(bulkDetailsEntryArr[length]);
                            x0.readTag();
                            length++;
                        }
                        bulkDetailsEntryArr[length] = new BulkDetailsEntry();
                        x0.readMessage(bulkDetailsEntryArr[length]);
                        this.entry = bulkDetailsEntryArr;
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

        public BulkDetailsResponse() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.entry != null && this.entry.length > 0) {
                for (int i = 0; i < this.entry.length; i++) {
                    BulkDetailsEntry element = this.entry[i];
                    if (element != null) {
                        output.writeMessage(1, element);
                    }
                }
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.entry != null && this.entry.length > 0) {
                for (int i = 0; i < this.entry.length; i++) {
                    BulkDetailsEntry element = this.entry[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(1, element);
                    }
                }
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class BulkDetailsEntry extends MessageNano {
        private static volatile BulkDetailsEntry[] _emptyArray;
        public DocV2 doc = null;

        public static BulkDetailsEntry[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new BulkDetailsEntry[0];
                    }
                }
            }
            return _emptyArray;
        }

        public BulkDetailsEntry() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.doc != null) {
                output.writeMessage(1, this.doc);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.doc != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(1, this.doc);
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
                    case 10:
                        if (this.doc == null) {
                            this.doc = new DocV2();
                        }
                        x0.readMessage(this.doc);
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

    /* loaded from: classes.dex */
    public static final class DiscoveryBadge extends MessageNano {
        private static volatile DiscoveryBadge[] _emptyArray;
        public String title = "";
        public boolean hasTitle = false;
        public String contentDescription = "";
        public boolean hasContentDescription = false;
        public Common.Image image = null;
        public int backgroundColor = 0;
        public boolean hasBackgroundColor = false;
        public DiscoveryBadgeLink discoveryBadgeLink = null;
        public byte[] serverLogsCookie = WireFormatNano.EMPTY_BYTES;
        public boolean hasServerLogsCookie = false;
        public boolean isPlusOne = false;
        public boolean hasIsPlusOne = false;
        public float aggregateRating = 0.0f;
        public boolean hasAggregateRating = false;
        public int userStarRating = 0;
        public boolean hasUserStarRating = false;
        public String downloadCount = "";
        public boolean hasDownloadCount = false;
        public String downloadUnits = "";
        public boolean hasDownloadUnits = false;
        public PlayerBadge playerBadge = null;
        public FamilyAgeRangeBadge familyAgeRangeBadge = null;
        public FamilyCategoryBadge familyCategoryBadge = null;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.title = x0.readString();
                        this.hasTitle = true;
                        break;
                    case 18:
                        if (this.image == null) {
                            this.image = new Common.Image();
                        }
                        x0.readMessage(this.image);
                        break;
                    case 24:
                        this.backgroundColor = x0.readRawVarint32();
                        this.hasBackgroundColor = true;
                        break;
                    case 34:
                        if (this.discoveryBadgeLink == null) {
                            this.discoveryBadgeLink = new DiscoveryBadgeLink();
                        }
                        x0.readMessage(this.discoveryBadgeLink);
                        break;
                    case 42:
                        this.serverLogsCookie = x0.readBytes();
                        this.hasServerLogsCookie = true;
                        break;
                    case 48:
                        this.isPlusOne = x0.readBool();
                        this.hasIsPlusOne = true;
                        break;
                    case 61:
                        this.aggregateRating = Float.intBitsToFloat(x0.readRawLittleEndian32());
                        this.hasAggregateRating = true;
                        break;
                    case 64:
                        this.userStarRating = x0.readRawVarint32();
                        this.hasUserStarRating = true;
                        break;
                    case 74:
                        this.downloadCount = x0.readString();
                        this.hasDownloadCount = true;
                        break;
                    case 82:
                        this.downloadUnits = x0.readString();
                        this.hasDownloadUnits = true;
                        break;
                    case 90:
                        this.contentDescription = x0.readString();
                        this.hasContentDescription = true;
                        break;
                    case 98:
                        if (this.playerBadge == null) {
                            this.playerBadge = new PlayerBadge();
                        }
                        x0.readMessage(this.playerBadge);
                        break;
                    case 106:
                        if (this.familyAgeRangeBadge == null) {
                            this.familyAgeRangeBadge = new FamilyAgeRangeBadge();
                        }
                        x0.readMessage(this.familyAgeRangeBadge);
                        break;
                    case 114:
                        if (this.familyCategoryBadge == null) {
                            this.familyCategoryBadge = new FamilyCategoryBadge();
                        }
                        x0.readMessage(this.familyCategoryBadge);
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

        public static DiscoveryBadge[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new DiscoveryBadge[0];
                    }
                }
            }
            return _emptyArray;
        }

        public DiscoveryBadge() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasTitle || !this.title.equals("")) {
                output.writeString(1, this.title);
            }
            if (this.image != null) {
                output.writeMessage(2, this.image);
            }
            if (this.hasBackgroundColor || this.backgroundColor != 0) {
                output.writeInt32(3, this.backgroundColor);
            }
            if (this.discoveryBadgeLink != null) {
                output.writeMessage(4, this.discoveryBadgeLink);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                output.writeBytes(5, this.serverLogsCookie);
            }
            if (this.hasIsPlusOne || this.isPlusOne) {
                output.writeBool(6, this.isPlusOne);
            }
            if (this.hasAggregateRating || Float.floatToIntBits(this.aggregateRating) != Float.floatToIntBits(0.0f)) {
                output.writeFloat(7, this.aggregateRating);
            }
            if (this.hasUserStarRating || this.userStarRating != 0) {
                output.writeInt32(8, this.userStarRating);
            }
            if (this.hasDownloadCount || !this.downloadCount.equals("")) {
                output.writeString(9, this.downloadCount);
            }
            if (this.hasDownloadUnits || !this.downloadUnits.equals("")) {
                output.writeString(10, this.downloadUnits);
            }
            if (this.hasContentDescription || !this.contentDescription.equals("")) {
                output.writeString(11, this.contentDescription);
            }
            if (this.playerBadge != null) {
                output.writeMessage(12, this.playerBadge);
            }
            if (this.familyAgeRangeBadge != null) {
                output.writeMessage(13, this.familyAgeRangeBadge);
            }
            if (this.familyCategoryBadge != null) {
                output.writeMessage(14, this.familyCategoryBadge);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasTitle || !this.title.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.title);
            }
            if (this.image != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.image);
            }
            if (this.hasBackgroundColor || this.backgroundColor != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(3, this.backgroundColor);
            }
            if (this.discoveryBadgeLink != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(4, this.discoveryBadgeLink);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                size += CodedOutputByteBufferNano.computeBytesSize(5, this.serverLogsCookie);
            }
            if (this.hasIsPlusOne || this.isPlusOne) {
                size += CodedOutputByteBufferNano.computeTagSize(6) + 1;
            }
            if (this.hasAggregateRating || Float.floatToIntBits(this.aggregateRating) != Float.floatToIntBits(0.0f)) {
                size += CodedOutputByteBufferNano.computeTagSize(7) + 4;
            }
            if (this.hasUserStarRating || this.userStarRating != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(8, this.userStarRating);
            }
            if (this.hasDownloadCount || !this.downloadCount.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(9, this.downloadCount);
            }
            if (this.hasDownloadUnits || !this.downloadUnits.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(10, this.downloadUnits);
            }
            if (this.hasContentDescription || !this.contentDescription.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(11, this.contentDescription);
            }
            if (this.playerBadge != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(12, this.playerBadge);
            }
            if (this.familyAgeRangeBadge != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(13, this.familyAgeRangeBadge);
            }
            if (this.familyCategoryBadge != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(14, this.familyCategoryBadge);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class PlayerBadge extends MessageNano {
        public Common.Image overlayIcon = null;

        public PlayerBadge() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.overlayIcon != null) {
                output.writeMessage(1, this.overlayIcon);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.overlayIcon != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(1, this.overlayIcon);
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
                    case 10:
                        if (this.overlayIcon == null) {
                            this.overlayIcon = new Common.Image();
                        }
                        x0.readMessage(this.overlayIcon);
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

    /* loaded from: classes.dex */
    public static final class FamilyAgeRangeBadge extends MessageNano {
        public FamilyAgeRangeBadge() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int readTag;
            do {
                readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        return this;
                }
            } while (WireFormatNano.parseUnknownField(x0, readTag));
            return this;
        }
    }

    /* loaded from: classes.dex */
    public static final class FamilyCategoryBadge extends MessageNano {
        public FamilyCategoryBadge() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int readTag;
            do {
                readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        return this;
                }
            } while (WireFormatNano.parseUnknownField(x0, readTag));
            return this;
        }
    }

    /* loaded from: classes.dex */
    public static final class DiscoveryBadgeLink extends MessageNano {
        public Link link = null;
        public String userReviewsUrl = "";
        public boolean hasUserReviewsUrl = false;
        public String criticReviewsUrl = "";
        public boolean hasCriticReviewsUrl = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        if (this.link == null) {
                            this.link = new Link();
                        }
                        x0.readMessage(this.link);
                        break;
                    case 18:
                        this.userReviewsUrl = x0.readString();
                        this.hasUserReviewsUrl = true;
                        break;
                    case 26:
                        this.criticReviewsUrl = x0.readString();
                        this.hasCriticReviewsUrl = true;
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

        public DiscoveryBadgeLink() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.link != null) {
                output.writeMessage(1, this.link);
            }
            if (this.hasUserReviewsUrl || !this.userReviewsUrl.equals("")) {
                output.writeString(2, this.userReviewsUrl);
            }
            if (this.hasCriticReviewsUrl || !this.criticReviewsUrl.equals("")) {
                output.writeString(3, this.criticReviewsUrl);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.link != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.link);
            }
            if (this.hasUserReviewsUrl || !this.userReviewsUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.userReviewsUrl);
            }
            if (this.hasCriticReviewsUrl || !this.criticReviewsUrl.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(3, this.criticReviewsUrl);
            }
            return size;
        }
    }
}
