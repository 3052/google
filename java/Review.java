package com.google.android.finsky.protos;

import com.google.android.finsky.protos.Common;
import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class Review extends MessageNano {
    private static volatile Review[] _emptyArray;
    public String commentId = "";
    public boolean hasCommentId = false;
    public DocV2 author = null;
    public int starRating = 0;
    public boolean hasStarRating = false;
    public Common.Image sentiment = null;
    public String title = "";
    public boolean hasTitle = false;
    public String comment = "";
    public boolean hasComment = false;
    public String url = "";
    public boolean hasUrl = false;
    public String source = "";
    public boolean hasSource = false;
    public String documentVersion = "";
    public boolean hasDocumentVersion = false;
    public long timestampMsec = 0;
    public boolean hasTimestampMsec = false;
    public String deviceName = "";
    public boolean hasDeviceName = false;
    public String replyText = "";
    public boolean hasReplyText = false;
    public long replyTimestampMsec = 0;
    public boolean hasReplyTimestampMsec = false;
    public int helpfulCount = 0;
    public boolean hasHelpfulCount = false;
    public long thumbsUpCount = 0;
    public boolean hasThumbsUpCount = false;
    public OBSOLETE_PlusProfile oBSOLETEPlusProfile = null;
    public String authorName = "";
    public boolean hasAuthorName = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 10:
                    this.authorName = x0.readString();
                    this.hasAuthorName = true;
                    break;
                case 18:
                    this.url = x0.readString();
                    this.hasUrl = true;
                    break;
                case 26:
                    this.source = x0.readString();
                    this.hasSource = true;
                    break;
                case 34:
                    this.documentVersion = x0.readString();
                    this.hasDocumentVersion = true;
                    break;
                case 40:
                    this.timestampMsec = x0.readRawVarint64();
                    this.hasTimestampMsec = true;
                    break;
                case 48:
                    this.starRating = x0.readRawVarint32();
                    this.hasStarRating = true;
                    break;
                case 58:
                    this.title = x0.readString();
                    this.hasTitle = true;
                    break;
                case 66:
                    this.comment = x0.readString();
                    this.hasComment = true;
                    break;
                case 74:
                    this.commentId = x0.readString();
                    this.hasCommentId = true;
                    break;
                case 154:
                    this.deviceName = x0.readString();
                    this.hasDeviceName = true;
                    break;
                case 234:
                    this.replyText = x0.readString();
                    this.hasReplyText = true;
                    break;
                case 240:
                    this.replyTimestampMsec = x0.readRawVarint64();
                    this.hasReplyTimestampMsec = true;
                    break;
                case 250:
                    if (this.oBSOLETEPlusProfile == null) {
                        this.oBSOLETEPlusProfile = new OBSOLETE_PlusProfile();
                    }
                    x0.readMessage(this.oBSOLETEPlusProfile);
                    break;
                case 266:
                    if (this.author == null) {
                        this.author = new DocV2();
                    }
                    x0.readMessage(this.author);
                    break;
                case 274:
                    if (this.sentiment == null) {
                        this.sentiment = new Common.Image();
                    }
                    x0.readMessage(this.sentiment);
                    break;
                case 280:
                    this.helpfulCount = x0.readRawVarint32();
                    this.hasHelpfulCount = true;
                    break;
                case 304:
                    this.thumbsUpCount = x0.readRawVarint64();
                    this.hasThumbsUpCount = true;
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

    public static Review[] emptyArray() {
        if (_emptyArray == null) {
            synchronized (InternalNano.LAZY_INIT_LOCK) {
                if (_emptyArray == null) {
                    _emptyArray = new Review[0];
                }
            }
        }
        return _emptyArray;
    }

    public Review() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasAuthorName || !this.authorName.equals("")) {
            output.writeString(1, this.authorName);
        }
        if (this.hasUrl || !this.url.equals("")) {
            output.writeString(2, this.url);
        }
        if (this.hasSource || !this.source.equals("")) {
            output.writeString(3, this.source);
        }
        if (this.hasDocumentVersion || !this.documentVersion.equals("")) {
            output.writeString(4, this.documentVersion);
        }
        if (this.hasTimestampMsec || this.timestampMsec != 0) {
            output.writeInt64(5, this.timestampMsec);
        }
        if (this.hasStarRating || this.starRating != 0) {
            output.writeInt32(6, this.starRating);
        }
        if (this.hasTitle || !this.title.equals("")) {
            output.writeString(7, this.title);
        }
        if (this.hasComment || !this.comment.equals("")) {
            output.writeString(8, this.comment);
        }
        if (this.hasCommentId || !this.commentId.equals("")) {
            output.writeString(9, this.commentId);
        }
        if (this.hasDeviceName || !this.deviceName.equals("")) {
            output.writeString(19, this.deviceName);
        }
        if (this.hasReplyText || !this.replyText.equals("")) {
            output.writeString(29, this.replyText);
        }
        if (this.hasReplyTimestampMsec || this.replyTimestampMsec != 0) {
            output.writeInt64(30, this.replyTimestampMsec);
        }
        if (this.oBSOLETEPlusProfile != null) {
            output.writeMessage(31, this.oBSOLETEPlusProfile);
        }
        if (this.author != null) {
            output.writeMessage(33, this.author);
        }
        if (this.sentiment != null) {
            output.writeMessage(34, this.sentiment);
        }
        if (this.hasHelpfulCount || this.helpfulCount != 0) {
            output.writeInt32(35, this.helpfulCount);
        }
        if (this.hasThumbsUpCount || this.thumbsUpCount != 0) {
            output.writeUInt64(38, this.thumbsUpCount);
        }
        super.writeTo(output);
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasAuthorName || !this.authorName.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(1, this.authorName);
        }
        if (this.hasUrl || !this.url.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(2, this.url);
        }
        if (this.hasSource || !this.source.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(3, this.source);
        }
        if (this.hasDocumentVersion || !this.documentVersion.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(4, this.documentVersion);
        }
        if (this.hasTimestampMsec || this.timestampMsec != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(5, this.timestampMsec);
        }
        if (this.hasStarRating || this.starRating != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(6, this.starRating);
        }
        if (this.hasTitle || !this.title.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(7, this.title);
        }
        if (this.hasComment || !this.comment.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(8, this.comment);
        }
        if (this.hasCommentId || !this.commentId.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(9, this.commentId);
        }
        if (this.hasDeviceName || !this.deviceName.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(19, this.deviceName);
        }
        if (this.hasReplyText || !this.replyText.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(29, this.replyText);
        }
        if (this.hasReplyTimestampMsec || this.replyTimestampMsec != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(30, this.replyTimestampMsec);
        }
        if (this.oBSOLETEPlusProfile != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(31, this.oBSOLETEPlusProfile);
        }
        if (this.author != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(33, this.author);
        }
        if (this.sentiment != null) {
            size += CodedOutputByteBufferNano.computeMessageSize(34, this.sentiment);
        }
        if (this.hasHelpfulCount || this.helpfulCount != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(35, this.helpfulCount);
        }
        if (this.hasThumbsUpCount || this.thumbsUpCount != 0) {
            return size + CodedOutputByteBufferNano.computeUInt64Size(38, this.thumbsUpCount);
        }
        return size;
    }
}
