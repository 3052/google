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
public interface SearchSuggest {

    /* loaded from: classes.dex */
    public static final class Suggestion extends MessageNano {
        private static volatile Suggestion[] _emptyArray;
        public int type = 1;
        public boolean hasType = false;
        public String displayText = "";
        public boolean hasDisplayText = false;
        public Common.Image image = null;
        public Link link = null;
        public byte[] serverLogsCookie = WireFormatNano.EMPTY_BYTES;
        public boolean hasServerLogsCookie = false;
        public DocV2 document = null;
        public int searchBackend = 0;
        public boolean hasSearchBackend = false;
        public String suggestedQuery = "";
        public boolean hasSuggestedQuery = false;
        public NavSuggestion navSuggestion = null;

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
                                this.type = readRawVarint32;
                                this.hasType = true;
                                continue;
                        }
                    case 18:
                        this.suggestedQuery = x0.readString();
                        this.hasSuggestedQuery = true;
                        break;
                    case 26:
                        if (this.navSuggestion == null) {
                            this.navSuggestion = new NavSuggestion();
                        }
                        x0.readMessage(this.navSuggestion);
                        break;
                    case 34:
                        this.serverLogsCookie = x0.readBytes();
                        this.hasServerLogsCookie = true;
                        break;
                    case 42:
                        if (this.image == null) {
                            this.image = new Common.Image();
                        }
                        x0.readMessage(this.image);
                        break;
                    case 50:
                        this.displayText = x0.readString();
                        this.hasDisplayText = true;
                        break;
                    case 58:
                        if (this.link == null) {
                            this.link = new Link();
                        }
                        x0.readMessage(this.link);
                        break;
                    case 66:
                        if (this.document == null) {
                            this.document = new DocV2();
                        }
                        x0.readMessage(this.document);
                        break;
                    case 72:
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
                                this.searchBackend = readRawVarint322;
                                this.hasSearchBackend = true;
                                continue;
                        }
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

        public static Suggestion[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new Suggestion[0];
                    }
                }
            }
            return _emptyArray;
        }

        public Suggestion() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.type != 1 || this.hasType) {
                output.writeInt32(1, this.type);
            }
            if (this.hasSuggestedQuery || !this.suggestedQuery.equals("")) {
                output.writeString(2, this.suggestedQuery);
            }
            if (this.navSuggestion != null) {
                output.writeMessage(3, this.navSuggestion);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                output.writeBytes(4, this.serverLogsCookie);
            }
            if (this.image != null) {
                output.writeMessage(5, this.image);
            }
            if (this.hasDisplayText || !this.displayText.equals("")) {
                output.writeString(6, this.displayText);
            }
            if (this.link != null) {
                output.writeMessage(7, this.link);
            }
            if (this.document != null) {
                output.writeMessage(8, this.document);
            }
            if (this.searchBackend != 0 || this.hasSearchBackend) {
                output.writeInt32(9, this.searchBackend);
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
            if (this.hasSuggestedQuery || !this.suggestedQuery.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.suggestedQuery);
            }
            if (this.navSuggestion != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.navSuggestion);
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                size += CodedOutputByteBufferNano.computeBytesSize(4, this.serverLogsCookie);
            }
            if (this.image != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(5, this.image);
            }
            if (this.hasDisplayText || !this.displayText.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(6, this.displayText);
            }
            if (this.link != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(7, this.link);
            }
            if (this.document != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(8, this.document);
            }
            if (this.searchBackend != 0 || this.hasSearchBackend) {
                return size + CodedOutputByteBufferNano.computeInt32Size(9, this.searchBackend);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class SearchSuggestResponse extends MessageNano {
        public Suggestion[] suggestion = Suggestion.emptyArray();
        public byte[] serverLogsCookie = WireFormatNano.EMPTY_BYTES;
        public boolean hasServerLogsCookie = false;

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
                        if (this.suggestion == null) {
                            length = 0;
                        } else {
                            length = this.suggestion.length;
                        }
                        Suggestion[] suggestionArr = new Suggestion[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.suggestion, 0, suggestionArr, 0, length);
                        }
                        while (length < suggestionArr.length - 1) {
                            suggestionArr[length] = new Suggestion();
                            x0.readMessage(suggestionArr[length]);
                            x0.readTag();
                            length++;
                        }
                        suggestionArr[length] = new Suggestion();
                        x0.readMessage(suggestionArr[length]);
                        this.suggestion = suggestionArr;
                        break;
                    case 18:
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

        public SearchSuggestResponse() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.suggestion != null && this.suggestion.length > 0) {
                for (int i = 0; i < this.suggestion.length; i++) {
                    Suggestion element = this.suggestion[i];
                    if (element != null) {
                        output.writeMessage(1, element);
                    }
                }
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                output.writeBytes(2, this.serverLogsCookie);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.suggestion != null && this.suggestion.length > 0) {
                for (int i = 0; i < this.suggestion.length; i++) {
                    Suggestion element = this.suggestion[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(1, element);
                    }
                }
            }
            if (this.hasServerLogsCookie || !Arrays.equals(this.serverLogsCookie, WireFormatNano.EMPTY_BYTES)) {
                return size + CodedOutputByteBufferNano.computeBytesSize(2, this.serverLogsCookie);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class NavSuggestion extends MessageNano {
        public String docId = "";
        public boolean hasDocId = false;
        public String description = "";
        public boolean hasDescription = false;
        public byte[] imageBlob = WireFormatNano.EMPTY_BYTES;
        public boolean hasImageBlob = false;
        public Common.Image image = null;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.docId = x0.readString();
                        this.hasDocId = true;
                        break;
                    case 18:
                        this.imageBlob = x0.readBytes();
                        this.hasImageBlob = true;
                        break;
                    case 26:
                        if (this.image == null) {
                            this.image = new Common.Image();
                        }
                        x0.readMessage(this.image);
                        break;
                    case 34:
                        this.description = x0.readString();
                        this.hasDescription = true;
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

        public NavSuggestion() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasDocId || !this.docId.equals("")) {
                output.writeString(1, this.docId);
            }
            if (this.hasImageBlob || !Arrays.equals(this.imageBlob, WireFormatNano.EMPTY_BYTES)) {
                output.writeBytes(2, this.imageBlob);
            }
            if (this.image != null) {
                output.writeMessage(3, this.image);
            }
            if (this.hasDescription || !this.description.equals("")) {
                output.writeString(4, this.description);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasDocId || !this.docId.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.docId);
            }
            if (this.hasImageBlob || !Arrays.equals(this.imageBlob, WireFormatNano.EMPTY_BYTES)) {
                size += CodedOutputByteBufferNano.computeBytesSize(2, this.imageBlob);
            }
            if (this.image != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.image);
            }
            if (this.hasDescription || !this.description.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(4, this.description);
            }
            return size;
        }
    }
}
