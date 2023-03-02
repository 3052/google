package com.google.android.finsky.protos;

import com.google.android.finsky.protos.Common;
import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public interface DocDetails {

    /* loaded from: classes.dex */
    public static final class DocumentDetails extends MessageNano {
        public AppDetails appDetails = null;
        public AlbumDetails albumDetails = null;
        public ArtistDetails artistDetails = null;
        public SongDetails songDetails = null;
        public BookDetails bookDetails = null;
        public VideoDetails videoDetails = null;
        public SubscriptionDetails subscriptionDetails = null;
        public MagazineDetails magazineDetails = null;
        public TvShowDetails tvShowDetails = null;
        public TvSeasonDetails tvSeasonDetails = null;
        public TvEpisodeDetails tvEpisodeDetails = null;
        public PersonDetails personDetails = null;
        public TalentDetails talentDetails = null;
        public DeveloperDetails developerDetails = null;
        public BookSeriesDetails bookSeriesDetails = null;

        public DocumentDetails() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.appDetails != null) {
                output.writeMessage(1, this.appDetails);
            }
            if (this.albumDetails != null) {
                output.writeMessage(2, this.albumDetails);
            }
            if (this.artistDetails != null) {
                output.writeMessage(3, this.artistDetails);
            }
            if (this.songDetails != null) {
                output.writeMessage(4, this.songDetails);
            }
            if (this.bookDetails != null) {
                output.writeMessage(5, this.bookDetails);
            }
            if (this.videoDetails != null) {
                output.writeMessage(6, this.videoDetails);
            }
            if (this.subscriptionDetails != null) {
                output.writeMessage(7, this.subscriptionDetails);
            }
            if (this.magazineDetails != null) {
                output.writeMessage(8, this.magazineDetails);
            }
            if (this.tvShowDetails != null) {
                output.writeMessage(9, this.tvShowDetails);
            }
            if (this.tvSeasonDetails != null) {
                output.writeMessage(10, this.tvSeasonDetails);
            }
            if (this.tvEpisodeDetails != null) {
                output.writeMessage(11, this.tvEpisodeDetails);
            }
            if (this.personDetails != null) {
                output.writeMessage(12, this.personDetails);
            }
            if (this.talentDetails != null) {
                output.writeMessage(13, this.talentDetails);
            }
            if (this.developerDetails != null) {
                output.writeMessage(14, this.developerDetails);
            }
            if (this.bookSeriesDetails != null) {
                output.writeMessage(15, this.bookSeriesDetails);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.appDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.appDetails);
            }
            if (this.albumDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.albumDetails);
            }
            if (this.artistDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.artistDetails);
            }
            if (this.songDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(4, this.songDetails);
            }
            if (this.bookDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(5, this.bookDetails);
            }
            if (this.videoDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(6, this.videoDetails);
            }
            if (this.subscriptionDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(7, this.subscriptionDetails);
            }
            if (this.magazineDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(8, this.magazineDetails);
            }
            if (this.tvShowDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(9, this.tvShowDetails);
            }
            if (this.tvSeasonDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(10, this.tvSeasonDetails);
            }
            if (this.tvEpisodeDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(11, this.tvEpisodeDetails);
            }
            if (this.personDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(12, this.personDetails);
            }
            if (this.talentDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(13, this.talentDetails);
            }
            if (this.developerDetails != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(14, this.developerDetails);
            }
            if (this.bookSeriesDetails != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(15, this.bookSeriesDetails);
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
                        if (this.appDetails == null) {
                            this.appDetails = new AppDetails();
                        }
                        x0.readMessage(this.appDetails);
                        break;
                    case 18:
                        if (this.albumDetails == null) {
                            this.albumDetails = new AlbumDetails();
                        }
                        x0.readMessage(this.albumDetails);
                        break;
                    case 26:
                        if (this.artistDetails == null) {
                            this.artistDetails = new ArtistDetails();
                        }
                        x0.readMessage(this.artistDetails);
                        break;
                    case 34:
                        if (this.songDetails == null) {
                            this.songDetails = new SongDetails();
                        }
                        x0.readMessage(this.songDetails);
                        break;
                    case 42:
                        if (this.bookDetails == null) {
                            this.bookDetails = new BookDetails();
                        }
                        x0.readMessage(this.bookDetails);
                        break;
                    case 50:
                        if (this.videoDetails == null) {
                            this.videoDetails = new VideoDetails();
                        }
                        x0.readMessage(this.videoDetails);
                        break;
                    case 58:
                        if (this.subscriptionDetails == null) {
                            this.subscriptionDetails = new SubscriptionDetails();
                        }
                        x0.readMessage(this.subscriptionDetails);
                        break;
                    case 66:
                        if (this.magazineDetails == null) {
                            this.magazineDetails = new MagazineDetails();
                        }
                        x0.readMessage(this.magazineDetails);
                        break;
                    case 74:
                        if (this.tvShowDetails == null) {
                            this.tvShowDetails = new TvShowDetails();
                        }
                        x0.readMessage(this.tvShowDetails);
                        break;
                    case 82:
                        if (this.tvSeasonDetails == null) {
                            this.tvSeasonDetails = new TvSeasonDetails();
                        }
                        x0.readMessage(this.tvSeasonDetails);
                        break;
                    case 90:
                        if (this.tvEpisodeDetails == null) {
                            this.tvEpisodeDetails = new TvEpisodeDetails();
                        }
                        x0.readMessage(this.tvEpisodeDetails);
                        break;
                    case 98:
                        if (this.personDetails == null) {
                            this.personDetails = new PersonDetails();
                        }
                        x0.readMessage(this.personDetails);
                        break;
                    case 106:
                        if (this.talentDetails == null) {
                            this.talentDetails = new TalentDetails();
                        }
                        x0.readMessage(this.talentDetails);
                        break;
                    case 114:
                        if (this.developerDetails == null) {
                            this.developerDetails = new DeveloperDetails();
                        }
                        x0.readMessage(this.developerDetails);
                        break;
                    case 122:
                        if (this.bookSeriesDetails == null) {
                            this.bookSeriesDetails = new BookSeriesDetails();
                        }
                        x0.readMessage(this.bookSeriesDetails);
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
    public static final class ProductDetails extends MessageNano {
        public String title = "";
        public boolean hasTitle = false;
        public ProductDetailsSection[] section = ProductDetailsSection.emptyArray();

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
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
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 18);
                        if (this.section == null) {
                            length = 0;
                        } else {
                            length = this.section.length;
                        }
                        ProductDetailsSection[] productDetailsSectionArr = new ProductDetailsSection[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.section, 0, productDetailsSectionArr, 0, length);
                        }
                        while (length < productDetailsSectionArr.length - 1) {
                            productDetailsSectionArr[length] = new ProductDetailsSection();
                            x0.readMessage(productDetailsSectionArr[length]);
                            x0.readTag();
                            length++;
                        }
                        productDetailsSectionArr[length] = new ProductDetailsSection();
                        x0.readMessage(productDetailsSectionArr[length]);
                        this.section = productDetailsSectionArr;
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

        public ProductDetails() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasTitle || !this.title.equals("")) {
                output.writeString(1, this.title);
            }
            if (this.section != null && this.section.length > 0) {
                for (int i = 0; i < this.section.length; i++) {
                    ProductDetailsSection element = this.section[i];
                    if (element != null) {
                        output.writeMessage(2, element);
                    }
                }
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasTitle || !this.title.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.title);
            }
            if (this.section != null && this.section.length > 0) {
                for (int i = 0; i < this.section.length; i++) {
                    ProductDetailsSection element = this.section[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(2, element);
                    }
                }
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class ProductDetailsSection extends MessageNano {
        private static volatile ProductDetailsSection[] _emptyArray;
        public String title = "";
        public boolean hasTitle = false;
        public ProductDetailsDescription[] description = ProductDetailsDescription.emptyArray();

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.title = x0.readString();
                        this.hasTitle = true;
                        break;
                    case 26:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 26);
                        if (this.description == null) {
                            length = 0;
                        } else {
                            length = this.description.length;
                        }
                        ProductDetailsDescription[] productDetailsDescriptionArr = new ProductDetailsDescription[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.description, 0, productDetailsDescriptionArr, 0, length);
                        }
                        while (length < productDetailsDescriptionArr.length - 1) {
                            productDetailsDescriptionArr[length] = new ProductDetailsDescription();
                            x0.readMessage(productDetailsDescriptionArr[length]);
                            x0.readTag();
                            length++;
                        }
                        productDetailsDescriptionArr[length] = new ProductDetailsDescription();
                        x0.readMessage(productDetailsDescriptionArr[length]);
                        this.description = productDetailsDescriptionArr;
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

        public static ProductDetailsSection[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new ProductDetailsSection[0];
                    }
                }
            }
            return _emptyArray;
        }

        public ProductDetailsSection() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasTitle || !this.title.equals("")) {
                output.writeString(1, this.title);
            }
            if (this.description != null && this.description.length > 0) {
                for (int i = 0; i < this.description.length; i++) {
                    ProductDetailsDescription element = this.description[i];
                    if (element != null) {
                        output.writeMessage(3, element);
                    }
                }
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasTitle || !this.title.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.title);
            }
            if (this.description != null && this.description.length > 0) {
                for (int i = 0; i < this.description.length; i++) {
                    ProductDetailsDescription element = this.description[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(3, element);
                    }
                }
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class ProductDetailsDescription extends MessageNano {
        private static volatile ProductDetailsDescription[] _emptyArray;
        public Common.Image image = null;
        public String description = "";
        public boolean hasDescription = false;

        public static ProductDetailsDescription[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new ProductDetailsDescription[0];
                    }
                }
            }
            return _emptyArray;
        }

        public ProductDetailsDescription() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.image != null) {
                output.writeMessage(1, this.image);
            }
            if (this.hasDescription || !this.description.equals("")) {
                output.writeString(2, this.description);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.image != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.image);
            }
            if (this.hasDescription || !this.description.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(2, this.description);
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
                        if (this.image == null) {
                            this.image = new Common.Image();
                        }
                        x0.readMessage(this.image);
                        break;
                    case 18:
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
    }

    /* loaded from: classes.dex */
    public static final class PersonDetails extends MessageNano {
        public boolean personIsRequester = false;
        public boolean hasPersonIsRequester = false;
        public boolean isGplusUser = true;
        public boolean hasIsGplusUser = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.personIsRequester = x0.readBool();
                        this.hasPersonIsRequester = true;
                        break;
                    case 16:
                        this.isGplusUser = x0.readBool();
                        this.hasIsGplusUser = true;
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

        public PersonDetails() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasPersonIsRequester || this.personIsRequester) {
                output.writeBool(1, this.personIsRequester);
            }
            if (this.hasIsGplusUser || !this.isGplusUser) {
                output.writeBool(2, this.isGplusUser);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasPersonIsRequester || this.personIsRequester) {
                size += CodedOutputByteBufferNano.computeTagSize(1) + 1;
            }
            if (this.hasIsGplusUser || !this.isGplusUser) {
                return size + CodedOutputByteBufferNano.computeTagSize(2) + 1;
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class TalentDetails extends MessageNano {
        public TalentExternalLinks externalLinks = null;
        public int primaryRoleId = 0;
        public boolean hasPrimaryRoleId = false;

        public TalentDetails() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.externalLinks != null) {
                output.writeMessage(1, this.externalLinks);
            }
            if (this.primaryRoleId != 0 || this.hasPrimaryRoleId) {
                output.writeInt32(2, this.primaryRoleId);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.externalLinks != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.externalLinks);
            }
            if (this.primaryRoleId != 0 || this.hasPrimaryRoleId) {
                return size + CodedOutputByteBufferNano.computeInt32Size(2, this.primaryRoleId);
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
                        if (this.externalLinks == null) {
                            this.externalLinks = new TalentExternalLinks();
                        }
                        x0.readMessage(this.externalLinks);
                        break;
                    case 16:
                        int readRawVarint32 = x0.readRawVarint32();
                        switch (readRawVarint32) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 5:
                            case 6:
                                this.primaryRoleId = readRawVarint32;
                                this.hasPrimaryRoleId = true;
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
    }

    /* loaded from: classes.dex */
    public static final class TalentExternalLinks extends MessageNano {
        public Link[] websiteUrl = Link.emptyArray();
        public Link googlePlusProfileUrl = null;
        public Link youtubeChannelUrl = null;

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
                        if (this.websiteUrl == null) {
                            length = 0;
                        } else {
                            length = this.websiteUrl.length;
                        }
                        Link[] linkArr = new Link[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.websiteUrl, 0, linkArr, 0, length);
                        }
                        while (length < linkArr.length - 1) {
                            linkArr[length] = new Link();
                            x0.readMessage(linkArr[length]);
                            x0.readTag();
                            length++;
                        }
                        linkArr[length] = new Link();
                        x0.readMessage(linkArr[length]);
                        this.websiteUrl = linkArr;
                        break;
                    case 18:
                        if (this.googlePlusProfileUrl == null) {
                            this.googlePlusProfileUrl = new Link();
                        }
                        x0.readMessage(this.googlePlusProfileUrl);
                        break;
                    case 26:
                        if (this.youtubeChannelUrl == null) {
                            this.youtubeChannelUrl = new Link();
                        }
                        x0.readMessage(this.youtubeChannelUrl);
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

        public TalentExternalLinks() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.websiteUrl != null && this.websiteUrl.length > 0) {
                for (int i = 0; i < this.websiteUrl.length; i++) {
                    Link element = this.websiteUrl[i];
                    if (element != null) {
                        output.writeMessage(1, element);
                    }
                }
            }
            if (this.googlePlusProfileUrl != null) {
                output.writeMessage(2, this.googlePlusProfileUrl);
            }
            if (this.youtubeChannelUrl != null) {
                output.writeMessage(3, this.youtubeChannelUrl);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.websiteUrl != null && this.websiteUrl.length > 0) {
                for (int i = 0; i < this.websiteUrl.length; i++) {
                    Link element = this.websiteUrl[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(1, element);
                    }
                }
            }
            if (this.googlePlusProfileUrl != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.googlePlusProfileUrl);
            }
            if (this.youtubeChannelUrl != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(3, this.youtubeChannelUrl);
            }
            return size;
        }
    }
}
