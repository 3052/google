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
public interface Containers {

    /* loaded from: classes.dex */
    public static final class ContainerMetadata extends MessageNano {
        public String browseUrl = "";
        public boolean hasBrowseUrl = false;
        public String nextPageUrl = "";
        public boolean hasNextPageUrl = false;
        public double relevance = 0.0d;
        public boolean hasRelevance = false;
        public long estimatedResults = 0;
        public boolean hasEstimatedResults = false;
        public String analyticsCookie = "";
        public boolean hasAnalyticsCookie = false;
        public boolean ordered = false;
        public boolean hasOrdered = false;
        public ContainerView[] containerView = ContainerView.emptyArray();
        public Common.Image leftIcon = null;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.browseUrl = x0.readString();
                        this.hasBrowseUrl = true;
                        break;
                    case 18:
                        this.nextPageUrl = x0.readString();
                        this.hasNextPageUrl = true;
                        break;
                    case 25:
                        this.relevance = Double.longBitsToDouble(x0.readRawLittleEndian64());
                        this.hasRelevance = true;
                        break;
                    case 32:
                        this.estimatedResults = x0.readRawVarint64();
                        this.hasEstimatedResults = true;
                        break;
                    case 42:
                        this.analyticsCookie = x0.readString();
                        this.hasAnalyticsCookie = true;
                        break;
                    case 48:
                        this.ordered = x0.readBool();
                        this.hasOrdered = true;
                        break;
                    case 58:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 58);
                        if (this.containerView == null) {
                            length = 0;
                        } else {
                            length = this.containerView.length;
                        }
                        ContainerView[] containerViewArr = new ContainerView[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.containerView, 0, containerViewArr, 0, length);
                        }
                        while (length < containerViewArr.length - 1) {
                            containerViewArr[length] = new ContainerView();
                            x0.readMessage(containerViewArr[length]);
                            x0.readTag();
                            length++;
                        }
                        containerViewArr[length] = new ContainerView();
                        x0.readMessage(containerViewArr[length]);
                        this.containerView = containerViewArr;
                        break;
                    case 66:
                        if (this.leftIcon == null) {
                            this.leftIcon = new Common.Image();
                        }
                        x0.readMessage(this.leftIcon);
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

        public ContainerMetadata() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasBrowseUrl || !this.browseUrl.equals("")) {
                output.writeString(1, this.browseUrl);
            }
            if (this.hasNextPageUrl || !this.nextPageUrl.equals("")) {
                output.writeString(2, this.nextPageUrl);
            }
            if (this.hasRelevance || Double.doubleToLongBits(this.relevance) != Double.doubleToLongBits(0.0d)) {
                output.writeDouble(3, this.relevance);
            }
            if (this.hasEstimatedResults || this.estimatedResults != 0) {
                output.writeInt64(4, this.estimatedResults);
            }
            if (this.hasAnalyticsCookie || !this.analyticsCookie.equals("")) {
                output.writeString(5, this.analyticsCookie);
            }
            if (this.hasOrdered || this.ordered) {
                output.writeBool(6, this.ordered);
            }
            if (this.containerView != null && this.containerView.length > 0) {
                for (int i = 0; i < this.containerView.length; i++) {
                    ContainerView element = this.containerView[i];
                    if (element != null) {
                        output.writeMessage(7, element);
                    }
                }
            }
            if (this.leftIcon != null) {
                output.writeMessage(8, this.leftIcon);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasBrowseUrl || !this.browseUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.browseUrl);
            }
            if (this.hasNextPageUrl || !this.nextPageUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.nextPageUrl);
            }
            if (this.hasRelevance || Double.doubleToLongBits(this.relevance) != Double.doubleToLongBits(0.0d)) {
                size += CodedOutputByteBufferNano.computeTagSize(3) + 8;
            }
            if (this.hasEstimatedResults || this.estimatedResults != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(4, this.estimatedResults);
            }
            if (this.hasAnalyticsCookie || !this.analyticsCookie.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(5, this.analyticsCookie);
            }
            if (this.hasOrdered || this.ordered) {
                size += CodedOutputByteBufferNano.computeTagSize(6) + 1;
            }
            if (this.containerView != null && this.containerView.length > 0) {
                for (int i = 0; i < this.containerView.length; i++) {
                    ContainerView element = this.containerView[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(7, element);
                    }
                }
            }
            if (this.leftIcon != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(8, this.leftIcon);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class ContainerView extends MessageNano {
        private static volatile ContainerView[] _emptyArray;
        public boolean selected = false;
        public boolean hasSelected = false;
        public String title = "";
        public boolean hasTitle = false;
        public String listUrl = "";
        public boolean hasListUrl = false;
        public byte[] serverLogsCookie = WireFormatNano.EMPTY_BYTES;
        public boolean hasServerLogsCookie = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.selected = x0.readBool();
                        this.hasSelected = true;
                        break;
                    case 18:
                        this.title = x0.readString();
                        this.hasTitle = true;
                        break;
                    case 26:
                        this.listUrl = x0.readString();
                        this.hasListUrl = true;
                        break;
                    case 34:
                        this.serverLogsCookie = x0.readBytes();
                        this.hasServerLogsCookie = true;
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

        public static ContainerView[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new ContainerView[0];
                    }
                }
            }
            return _emptyArray;
        }

        public ContainerView() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasSelected || this.selected) {
                output.writeBool(1, this.selected);
            }
            if (this.hasTitle || !this.title.equals("")) {
                output.writeString(2, this.title);
            }
            if (this.hasListUrl || !this.listUrl.equals("")) {
                output.writeString(3, this.listUrl);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                output.writeBytes(4, this.serverLogsCookie);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasSelected || this.selected) {
                size += CodedOutputByteBufferNano.computeTagSize(1) + 1;
            }
            if (this.hasTitle || !this.title.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.title);
            }
            if (this.hasListUrl || !this.listUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(3, this.listUrl);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                return size + CodedOutputByteBufferNano.computeBytesSize(4, this.serverLogsCookie);
            }
            return size;
        }
    }
}
