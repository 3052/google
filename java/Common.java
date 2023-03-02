package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InternalNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public interface Common {

    /* loaded from: classes.dex */
    public static final class Docid extends MessageNano {
        private static volatile Docid[] _emptyArray;
        public String backendDocid = "";
        public boolean hasBackendDocid = false;
        public int type = 1;
        public boolean hasType = false;
        public int backend = 0;
        public boolean hasBackend = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.backendDocid = x0.readString();
                        this.hasBackendDocid = true;
                        break;
                    case 16:
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
                                this.type = readRawVarint32;
                                this.hasType = true;
                                continue;
                        }
                    case 24:
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
                                this.backend = readRawVarint322;
                                this.hasBackend = true;
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

        public static Docid[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new Docid[0];
                    }
                }
            }
            return _emptyArray;
        }

        public Docid() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasBackendDocid || !this.backendDocid.equals("")) {
                output.writeString(1, this.backendDocid);
            }
            if (this.type != 1 || this.hasType) {
                output.writeInt32(2, this.type);
            }
            if (this.backend != 0 || this.hasBackend) {
                output.writeInt32(3, this.backend);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasBackendDocid || !this.backendDocid.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.backendDocid);
            }
            if (this.type != 1 || this.hasType) {
                size += CodedOutputByteBufferNano.computeInt32Size(2, this.type);
            }
            if (this.backend != 0 || this.hasBackend) {
                return size + CodedOutputByteBufferNano.computeInt32Size(3, this.backend);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class OfferPaymentPeriod extends MessageNano {
        public TimePeriod duration = null;
        public MonthAndDay start = null;
        public MonthAndDay end = null;

        public OfferPaymentPeriod() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.duration != null) {
                output.writeMessage(1, this.duration);
            }
            if (this.start != null) {
                output.writeMessage(2, this.start);
            }
            if (this.end != null) {
                output.writeMessage(3, this.end);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.duration != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.duration);
            }
            if (this.start != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.start);
            }
            if (this.end != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(3, this.end);
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
                        if (this.duration == null) {
                            this.duration = new TimePeriod();
                        }
                        x0.readMessage(this.duration);
                        break;
                    case 18:
                        if (this.start == null) {
                            this.start = new MonthAndDay();
                        }
                        x0.readMessage(this.start);
                        break;
                    case 26:
                        if (this.end == null) {
                            this.end = new MonthAndDay();
                        }
                        x0.readMessage(this.end);
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
    public static final class OfferPaymentOverride extends MessageNano {
        private static volatile OfferPaymentOverride[] _emptyArray;
        public long micros = 0;
        public boolean hasMicros = false;
        public MonthAndDay start = null;
        public MonthAndDay end = null;

        public static OfferPaymentOverride[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new OfferPaymentOverride[0];
                    }
                }
            }
            return _emptyArray;
        }

        public OfferPaymentOverride() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasMicros || this.micros != 0) {
                output.writeInt64(1, this.micros);
            }
            if (this.start != null) {
                output.writeMessage(2, this.start);
            }
            if (this.end != null) {
                output.writeMessage(3, this.end);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasMicros || this.micros != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(1, this.micros);
            }
            if (this.start != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.start);
            }
            if (this.end != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(3, this.end);
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
                    case 8:
                        this.micros = x0.readRawVarint64();
                        this.hasMicros = true;
                        break;
                    case 18:
                        if (this.start == null) {
                            this.start = new MonthAndDay();
                        }
                        x0.readMessage(this.start);
                        break;
                    case 26:
                        if (this.end == null) {
                            this.end = new MonthAndDay();
                        }
                        x0.readMessage(this.end);
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
    public static final class OfferPayment extends MessageNano {
        private static volatile OfferPayment[] _emptyArray;
        public long micros = 0;
        public boolean hasMicros = false;
        public String currencyCode = "";
        public boolean hasCurrencyCode = false;
        public OfferPaymentPeriod offerPaymentPeriod = null;
        public OfferPaymentOverride[] offerPaymentOverride = OfferPaymentOverride.emptyArray();

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.micros = x0.readRawVarint64();
                        this.hasMicros = true;
                        break;
                    case 18:
                        this.currencyCode = x0.readString();
                        this.hasCurrencyCode = true;
                        break;
                    case 26:
                        if (this.offerPaymentPeriod == null) {
                            this.offerPaymentPeriod = new OfferPaymentPeriod();
                        }
                        x0.readMessage(this.offerPaymentPeriod);
                        break;
                    case 34:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 34);
                        if (this.offerPaymentOverride == null) {
                            length = 0;
                        } else {
                            length = this.offerPaymentOverride.length;
                        }
                        OfferPaymentOverride[] offerPaymentOverrideArr = new OfferPaymentOverride[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.offerPaymentOverride, 0, offerPaymentOverrideArr, 0, length);
                        }
                        while (length < offerPaymentOverrideArr.length - 1) {
                            offerPaymentOverrideArr[length] = new OfferPaymentOverride();
                            x0.readMessage(offerPaymentOverrideArr[length]);
                            x0.readTag();
                            length++;
                        }
                        offerPaymentOverrideArr[length] = new OfferPaymentOverride();
                        x0.readMessage(offerPaymentOverrideArr[length]);
                        this.offerPaymentOverride = offerPaymentOverrideArr;
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

        public static OfferPayment[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new OfferPayment[0];
                    }
                }
            }
            return _emptyArray;
        }

        public OfferPayment() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasMicros || this.micros != 0) {
                output.writeInt64(1, this.micros);
            }
            if (this.hasCurrencyCode || !this.currencyCode.equals("")) {
                output.writeString(2, this.currencyCode);
            }
            if (this.offerPaymentPeriod != null) {
                output.writeMessage(3, this.offerPaymentPeriod);
            }
            if (this.offerPaymentOverride != null && this.offerPaymentOverride.length > 0) {
                for (int i = 0; i < this.offerPaymentOverride.length; i++) {
                    OfferPaymentOverride element = this.offerPaymentOverride[i];
                    if (element != null) {
                        output.writeMessage(4, element);
                    }
                }
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasMicros || this.micros != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(1, this.micros);
            }
            if (this.hasCurrencyCode || !this.currencyCode.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.currencyCode);
            }
            if (this.offerPaymentPeriod != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.offerPaymentPeriod);
            }
            if (this.offerPaymentOverride != null && this.offerPaymentOverride.length > 0) {
                for (int i = 0; i < this.offerPaymentOverride.length; i++) {
                    OfferPaymentOverride element = this.offerPaymentOverride[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(4, element);
                    }
                }
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class Offer extends MessageNano {
        private static volatile Offer[] _emptyArray;
        public OfferPayment[] offerPayment = OfferPayment.emptyArray();
        public boolean repeatLastPayment = false;
        public boolean hasRepeatLastPayment = false;
        public long micros = 0;
        public boolean hasMicros = false;
        public String currencyCode = "";
        public boolean hasCurrencyCode = false;
        public String formattedAmount = "";
        public boolean hasFormattedAmount = false;
        public String formattedName = "";
        public boolean hasFormattedName = false;
        public String formattedDescription = "";
        public boolean hasFormattedDescription = false;
        public String buyButtonLabel = "";
        public boolean hasBuyButtonLabel = false;
        public boolean instantPurchaseEnabled = false;
        public boolean hasInstantPurchaseEnabled = false;
        public long fullPriceMicros = 0;
        public boolean hasFullPriceMicros = false;
        public String formattedFullAmount = "";
        public boolean hasFormattedFullAmount = false;
        public Offer[] convertedPrice = emptyArray();
        public boolean checkoutFlowRequired = false;
        public boolean hasCheckoutFlowRequired = false;
        public boolean temporarilyFree = false;
        public boolean hasTemporarilyFree = false;
        public int offerType = 1;
        public boolean hasOfferType = false;
        public int licensedOfferType = 1;
        public boolean hasLicensedOfferType = false;
        public LicenseTerms licenseTerms = null;
        public RentalTerms rentalTerms = null;
        public SubscriptionTerms subscriptionTerms = null;
        public SubscriptionContentTerms subscriptionContentTerms = null;
        public VoucherOfferTerms voucherTerms = null;
        public boolean preorder = false;
        public boolean hasPreorder = false;
        public long preorderFulfillmentDisplayDate = 0;
        public boolean hasPreorderFulfillmentDisplayDate = false;
        public long onSaleDate = 0;
        public boolean hasOnSaleDate = false;
        public int onSaleDateDisplayTimeZoneOffsetMsec = 0;
        public boolean hasOnSaleDateDisplayTimeZoneOffsetMsec = false;
        public String[] promotionLabel = WireFormatNano.EMPTY_STRING_ARRAY;
        public String offerId = "";
        public boolean hasOfferId = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            int length2;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.micros = x0.readRawVarint64();
                        this.hasMicros = true;
                        break;
                    case 18:
                        this.currencyCode = x0.readString();
                        this.hasCurrencyCode = true;
                        break;
                    case 26:
                        this.formattedAmount = x0.readString();
                        this.hasFormattedAmount = true;
                        break;
                    case 34:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 34);
                        if (this.convertedPrice == null) {
                            length2 = 0;
                        } else {
                            length2 = this.convertedPrice.length;
                        }
                        Offer[] offerArr = new Offer[repeatedFieldArrayLength + length2];
                        if (length2 != 0) {
                            System.arraycopy(this.convertedPrice, 0, offerArr, 0, length2);
                        }
                        while (length2 < offerArr.length - 1) {
                            offerArr[length2] = new Offer();
                            x0.readMessage(offerArr[length2]);
                            x0.readTag();
                            length2++;
                        }
                        offerArr[length2] = new Offer();
                        x0.readMessage(offerArr[length2]);
                        this.convertedPrice = offerArr;
                        break;
                    case 40:
                        this.checkoutFlowRequired = x0.readBool();
                        this.hasCheckoutFlowRequired = true;
                        break;
                    case 48:
                        this.fullPriceMicros = x0.readRawVarint64();
                        this.hasFullPriceMicros = true;
                        break;
                    case 58:
                        this.formattedFullAmount = x0.readString();
                        this.hasFormattedFullAmount = true;
                        break;
                    case 64:
                        int readRawVarint32 = x0.readRawVarint32();
                        switch (readRawVarint32) {
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 6:
                            case 7:
                            case 8:
                            case 9:
                            case 10:
                            case 11:
                            case 12:
                            case 13:
                                this.offerType = readRawVarint32;
                                this.hasOfferType = true;
                                continue;
                        }
                    case 74:
                        if (this.rentalTerms == null) {
                            this.rentalTerms = new RentalTerms();
                        }
                        x0.readMessage(this.rentalTerms);
                        break;
                    case 80:
                        this.onSaleDate = x0.readRawVarint64();
                        this.hasOnSaleDate = true;
                        break;
                    case 90:
                        int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 90);
                        int length3 = this.promotionLabel == null ? 0 : this.promotionLabel.length;
                        String[] strArr = new String[repeatedFieldArrayLength2 + length3];
                        if (length3 != 0) {
                            System.arraycopy(this.promotionLabel, 0, strArr, 0, length3);
                        }
                        while (length3 < strArr.length - 1) {
                            strArr[length3] = x0.readString();
                            x0.readTag();
                            length3++;
                        }
                        strArr[length3] = x0.readString();
                        this.promotionLabel = strArr;
                        break;
                    case 98:
                        if (this.subscriptionTerms == null) {
                            this.subscriptionTerms = new SubscriptionTerms();
                        }
                        x0.readMessage(this.subscriptionTerms);
                        break;
                    case 106:
                        this.formattedName = x0.readString();
                        this.hasFormattedName = true;
                        break;
                    case 114:
                        this.formattedDescription = x0.readString();
                        this.hasFormattedDescription = true;
                        break;
                    case 120:
                        this.preorder = x0.readBool();
                        this.hasPreorder = true;
                        break;
                    case 128:
                        this.onSaleDateDisplayTimeZoneOffsetMsec = x0.readRawVarint32();
                        this.hasOnSaleDateDisplayTimeZoneOffsetMsec = true;
                        break;
                    case 136:
                        int readRawVarint322 = x0.readRawVarint32();
                        switch (readRawVarint322) {
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 6:
                            case 7:
                            case 8:
                            case 9:
                            case 10:
                            case 11:
                            case 12:
                            case 13:
                                this.licensedOfferType = readRawVarint322;
                                this.hasLicensedOfferType = true;
                                continue;
                        }
                    case 146:
                        if (this.subscriptionContentTerms == null) {
                            this.subscriptionContentTerms = new SubscriptionContentTerms();
                        }
                        x0.readMessage(this.subscriptionContentTerms);
                        break;
                    case 154:
                        this.offerId = x0.readString();
                        this.hasOfferId = true;
                        break;
                    case 160:
                        this.preorderFulfillmentDisplayDate = x0.readRawVarint64();
                        this.hasPreorderFulfillmentDisplayDate = true;
                        break;
                    case 170:
                        if (this.licenseTerms == null) {
                            this.licenseTerms = new LicenseTerms();
                        }
                        x0.readMessage(this.licenseTerms);
                        break;
                    case 176:
                        this.temporarilyFree = x0.readBool();
                        this.hasTemporarilyFree = true;
                        break;
                    case 186:
                        if (this.voucherTerms == null) {
                            this.voucherTerms = new VoucherOfferTerms();
                        }
                        x0.readMessage(this.voucherTerms);
                        break;
                    case 194:
                        int repeatedFieldArrayLength3 = WireFormatNano.getRepeatedFieldArrayLength(x0, 194);
                        if (this.offerPayment == null) {
                            length = 0;
                        } else {
                            length = this.offerPayment.length;
                        }
                        OfferPayment[] offerPaymentArr = new OfferPayment[repeatedFieldArrayLength3 + length];
                        if (length != 0) {
                            System.arraycopy(this.offerPayment, 0, offerPaymentArr, 0, length);
                        }
                        while (length < offerPaymentArr.length - 1) {
                            offerPaymentArr[length] = new OfferPayment();
                            x0.readMessage(offerPaymentArr[length]);
                            x0.readTag();
                            length++;
                        }
                        offerPaymentArr[length] = new OfferPayment();
                        x0.readMessage(offerPaymentArr[length]);
                        this.offerPayment = offerPaymentArr;
                        break;
                    case 200:
                        this.repeatLastPayment = x0.readBool();
                        this.hasRepeatLastPayment = true;
                        break;
                    case 210:
                        this.buyButtonLabel = x0.readString();
                        this.hasBuyButtonLabel = true;
                        break;
                    case 216:
                        this.instantPurchaseEnabled = x0.readBool();
                        this.hasInstantPurchaseEnabled = true;
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

        public static Offer[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new Offer[0];
                    }
                }
            }
            return _emptyArray;
        }

        public Offer() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasMicros || this.micros != 0) {
                output.writeInt64(1, this.micros);
            }
            if (this.hasCurrencyCode || !this.currencyCode.equals("")) {
                output.writeString(2, this.currencyCode);
            }
            if (this.hasFormattedAmount || !this.formattedAmount.equals("")) {
                output.writeString(3, this.formattedAmount);
            }
            if (this.convertedPrice != null && this.convertedPrice.length > 0) {
                for (int i = 0; i < this.convertedPrice.length; i++) {
                    Offer element = this.convertedPrice[i];
                    if (element != null) {
                        output.writeMessage(4, element);
                    }
                }
            }
            if (this.hasCheckoutFlowRequired || this.checkoutFlowRequired) {
                output.writeBool(5, this.checkoutFlowRequired);
            }
            if (this.hasFullPriceMicros || this.fullPriceMicros != 0) {
                output.writeInt64(6, this.fullPriceMicros);
            }
            if (this.hasFormattedFullAmount || !this.formattedFullAmount.equals("")) {
                output.writeString(7, this.formattedFullAmount);
            }
            if (this.offerType != 1 || this.hasOfferType) {
                output.writeInt32(8, this.offerType);
            }
            if (this.rentalTerms != null) {
                output.writeMessage(9, this.rentalTerms);
            }
            if (this.hasOnSaleDate || this.onSaleDate != 0) {
                output.writeInt64(10, this.onSaleDate);
            }
            if (this.promotionLabel != null && this.promotionLabel.length > 0) {
                for (int i2 = 0; i2 < this.promotionLabel.length; i2++) {
                    String element2 = this.promotionLabel[i2];
                    if (element2 != null) {
                        output.writeString(11, element2);
                    }
                }
            }
            if (this.subscriptionTerms != null) {
                output.writeMessage(12, this.subscriptionTerms);
            }
            if (this.hasFormattedName || !this.formattedName.equals("")) {
                output.writeString(13, this.formattedName);
            }
            if (this.hasFormattedDescription || !this.formattedDescription.equals("")) {
                output.writeString(14, this.formattedDescription);
            }
            if (this.hasPreorder || this.preorder) {
                output.writeBool(15, this.preorder);
            }
            if (this.hasOnSaleDateDisplayTimeZoneOffsetMsec || this.onSaleDateDisplayTimeZoneOffsetMsec != 0) {
                output.writeInt32(16, this.onSaleDateDisplayTimeZoneOffsetMsec);
            }
            if (this.licensedOfferType != 1 || this.hasLicensedOfferType) {
                output.writeInt32(17, this.licensedOfferType);
            }
            if (this.subscriptionContentTerms != null) {
                output.writeMessage(18, this.subscriptionContentTerms);
            }
            if (this.hasOfferId || !this.offerId.equals("")) {
                output.writeString(19, this.offerId);
            }
            if (this.hasPreorderFulfillmentDisplayDate || this.preorderFulfillmentDisplayDate != 0) {
                output.writeInt64(20, this.preorderFulfillmentDisplayDate);
            }
            if (this.licenseTerms != null) {
                output.writeMessage(21, this.licenseTerms);
            }
            if (this.hasTemporarilyFree || this.temporarilyFree) {
                output.writeBool(22, this.temporarilyFree);
            }
            if (this.voucherTerms != null) {
                output.writeMessage(23, this.voucherTerms);
            }
            if (this.offerPayment != null && this.offerPayment.length > 0) {
                for (int i3 = 0; i3 < this.offerPayment.length; i3++) {
                    OfferPayment element3 = this.offerPayment[i3];
                    if (element3 != null) {
                        output.writeMessage(24, element3);
                    }
                }
            }
            if (this.hasRepeatLastPayment || this.repeatLastPayment) {
                output.writeBool(25, this.repeatLastPayment);
            }
            if (this.hasBuyButtonLabel || !this.buyButtonLabel.equals("")) {
                output.writeString(26, this.buyButtonLabel);
            }
            if (this.hasInstantPurchaseEnabled || this.instantPurchaseEnabled) {
                output.writeBool(27, this.instantPurchaseEnabled);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasMicros || this.micros != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(1, this.micros);
            }
            if (this.hasCurrencyCode || !this.currencyCode.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.currencyCode);
            }
            if (this.hasFormattedAmount || !this.formattedAmount.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(3, this.formattedAmount);
            }
            if (this.convertedPrice != null && this.convertedPrice.length > 0) {
                for (int i = 0; i < this.convertedPrice.length; i++) {
                    Offer element = this.convertedPrice[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(4, element);
                    }
                }
            }
            if (this.hasCheckoutFlowRequired || this.checkoutFlowRequired) {
                size += CodedOutputByteBufferNano.computeTagSize(5) + 1;
            }
            if (this.hasFullPriceMicros || this.fullPriceMicros != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(6, this.fullPriceMicros);
            }
            if (this.hasFormattedFullAmount || !this.formattedFullAmount.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(7, this.formattedFullAmount);
            }
            if (this.offerType != 1 || this.hasOfferType) {
                size += CodedOutputByteBufferNano.computeInt32Size(8, this.offerType);
            }
            if (this.rentalTerms != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(9, this.rentalTerms);
            }
            if (this.hasOnSaleDate || this.onSaleDate != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(10, this.onSaleDate);
            }
            if (this.promotionLabel != null && this.promotionLabel.length > 0) {
                int dataCount = 0;
                int dataSize = 0;
                for (int i2 = 0; i2 < this.promotionLabel.length; i2++) {
                    String element2 = this.promotionLabel[i2];
                    if (element2 != null) {
                        dataCount++;
                        dataSize += CodedOutputByteBufferNano.computeStringSizeNoTag(element2);
                    }
                }
                size = size + dataSize + (dataCount * 1);
            }
            if (this.subscriptionTerms != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(12, this.subscriptionTerms);
            }
            if (this.hasFormattedName || !this.formattedName.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(13, this.formattedName);
            }
            if (this.hasFormattedDescription || !this.formattedDescription.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(14, this.formattedDescription);
            }
            if (this.hasPreorder || this.preorder) {
                size += CodedOutputByteBufferNano.computeTagSize(15) + 1;
            }
            if (this.hasOnSaleDateDisplayTimeZoneOffsetMsec || this.onSaleDateDisplayTimeZoneOffsetMsec != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(16, this.onSaleDateDisplayTimeZoneOffsetMsec);
            }
            if (this.licensedOfferType != 1 || this.hasLicensedOfferType) {
                size += CodedOutputByteBufferNano.computeInt32Size(17, this.licensedOfferType);
            }
            if (this.subscriptionContentTerms != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(18, this.subscriptionContentTerms);
            }
            if (this.hasOfferId || !this.offerId.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(19, this.offerId);
            }
            if (this.hasPreorderFulfillmentDisplayDate || this.preorderFulfillmentDisplayDate != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(20, this.preorderFulfillmentDisplayDate);
            }
            if (this.licenseTerms != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(21, this.licenseTerms);
            }
            if (this.hasTemporarilyFree || this.temporarilyFree) {
                size += CodedOutputByteBufferNano.computeTagSize(22) + 1;
            }
            if (this.voucherTerms != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(23, this.voucherTerms);
            }
            if (this.offerPayment != null && this.offerPayment.length > 0) {
                for (int i3 = 0; i3 < this.offerPayment.length; i3++) {
                    OfferPayment element3 = this.offerPayment[i3];
                    if (element3 != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(24, element3);
                    }
                }
            }
            if (this.hasRepeatLastPayment || this.repeatLastPayment) {
                size += CodedOutputByteBufferNano.computeTagSize(25) + 1;
            }
            if (this.hasBuyButtonLabel || !this.buyButtonLabel.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(26, this.buyButtonLabel);
            }
            if (this.hasInstantPurchaseEnabled || this.instantPurchaseEnabled) {
                return size + CodedOutputByteBufferNano.computeTagSize(27) + 1;
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class Image extends MessageNano {
        private static volatile Image[] _emptyArray;
        public int imageType = 0;
        public boolean hasImageType = false;
        public int positionInSequence = 0;
        public boolean hasPositionInSequence = false;
        public Dimension dimension = null;
        public String imageUrl = "";
        public boolean hasImageUrl = false;
        public String secureUrl = "";
        public boolean hasSecureUrl = false;
        public String altTextLocalized = "";
        public boolean hasAltTextLocalized = false;
        public boolean supportsFifeUrlOptions = false;
        public boolean hasSupportsFifeUrlOptions = false;
        public boolean supportsFifeMonogramOption = false;
        public boolean hasSupportsFifeMonogramOption = false;
        public int durationSeconds = 0;
        public boolean hasDurationSeconds = false;
        public String fillColorRgb = "";
        public boolean hasFillColorRgb = false;
        public boolean autogen = false;
        public boolean hasAutogen = false;
        public Attribution attribution = null;
        public String backgroundColorRgb = "";
        public boolean hasBackgroundColorRgb = false;
        public ImagePalette palette = null;
        public Citation citation = null;
        public int deviceClass = 0;
        public boolean hasDeviceClass = false;

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
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 5:
                            case 6:
                            case 7:
                            case 8:
                            case 9:
                            case 13:
                            case 14:
                            case 15:
                                this.imageType = readRawVarint32;
                                this.hasImageType = true;
                                continue;
                        }
                    case 19:
                        if (this.dimension == null) {
                            this.dimension = new Dimension();
                        }
                        x0.readGroup(this.dimension, 2);
                        break;
                    case 42:
                        this.imageUrl = x0.readString();
                        this.hasImageUrl = true;
                        break;
                    case 50:
                        this.altTextLocalized = x0.readString();
                        this.hasAltTextLocalized = true;
                        break;
                    case 58:
                        this.secureUrl = x0.readString();
                        this.hasSecureUrl = true;
                        break;
                    case 64:
                        this.positionInSequence = x0.readRawVarint32();
                        this.hasPositionInSequence = true;
                        break;
                    case 72:
                        this.supportsFifeUrlOptions = x0.readBool();
                        this.hasSupportsFifeUrlOptions = true;
                        break;
                    case 83:
                        if (this.citation == null) {
                            this.citation = new Citation();
                        }
                        x0.readGroup(this.citation, 10);
                        break;
                    case 112:
                        this.durationSeconds = x0.readRawVarint32();
                        this.hasDurationSeconds = true;
                        break;
                    case 122:
                        this.fillColorRgb = x0.readString();
                        this.hasFillColorRgb = true;
                        break;
                    case 128:
                        this.autogen = x0.readBool();
                        this.hasAutogen = true;
                        break;
                    case 138:
                        if (this.attribution == null) {
                            this.attribution = new Attribution();
                        }
                        x0.readMessage(this.attribution);
                        break;
                    case 154:
                        this.backgroundColorRgb = x0.readString();
                        this.hasBackgroundColorRgb = true;
                        break;
                    case 162:
                        if (this.palette == null) {
                            this.palette = new ImagePalette();
                        }
                        x0.readMessage(this.palette);
                        break;
                    case 168:
                        int readRawVarint322 = x0.readRawVarint32();
                        switch (readRawVarint322) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 5:
                            case 6:
                                this.deviceClass = readRawVarint322;
                                this.hasDeviceClass = true;
                                continue;
                        }
                    case 176:
                        this.supportsFifeMonogramOption = x0.readBool();
                        this.hasSupportsFifeMonogramOption = true;
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

        /* loaded from: classes.dex */
        public static final class Dimension extends MessageNano {
            public int width = 0;
            public boolean hasWidth = false;
            public int height = 0;
            public boolean hasHeight = false;
            public int aspectRatio = 0;
            public boolean hasAspectRatio = false;

            @Override // com.google.protobuf.nano.MessageNano
            public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
                while (true) {
                    int readTag = x0.readTag();
                    switch (readTag) {
                        case 0:
                            break;
                        case 24:
                            this.width = x0.readRawVarint32();
                            this.hasWidth = true;
                            break;
                        case 32:
                            this.height = x0.readRawVarint32();
                            this.hasHeight = true;
                            break;
                        case 144:
                            int readRawVarint32 = x0.readRawVarint32();
                            switch (readRawVarint32) {
                                case 0:
                                case 1:
                                case 2:
                                case 3:
                                case 4:
                                case 5:
                                case 6:
                                case 7:
                                    this.aspectRatio = readRawVarint32;
                                    this.hasAspectRatio = true;
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

            public Dimension() {
                this.cachedSize = -1;
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
                if (this.hasWidth || this.width != 0) {
                    output.writeInt32(3, this.width);
                }
                if (this.hasHeight || this.height != 0) {
                    output.writeInt32(4, this.height);
                }
                if (this.aspectRatio != 0 || this.hasAspectRatio) {
                    output.writeInt32(18, this.aspectRatio);
                }
                super.writeTo(output);
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final int computeSerializedSize() {
                int size = super.computeSerializedSize();
                if (this.hasWidth || this.width != 0) {
                    size += CodedOutputByteBufferNano.computeInt32Size(3, this.width);
                }
                if (this.hasHeight || this.height != 0) {
                    size += CodedOutputByteBufferNano.computeInt32Size(4, this.height);
                }
                if (this.aspectRatio != 0 || this.hasAspectRatio) {
                    return size + CodedOutputByteBufferNano.computeInt32Size(18, this.aspectRatio);
                }
                return size;
            }
        }

        /* loaded from: classes.dex */
        public static final class Citation extends MessageNano {
            public String titleLocalized = "";
            public boolean hasTitleLocalized = false;
            public String url = "";
            public boolean hasUrl = false;

            @Override // com.google.protobuf.nano.MessageNano
            public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
                while (true) {
                    int readTag = x0.readTag();
                    switch (readTag) {
                        case 0:
                            break;
                        case 90:
                            this.titleLocalized = x0.readString();
                            this.hasTitleLocalized = true;
                            break;
                        case 98:
                            this.url = x0.readString();
                            this.hasUrl = true;
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

            public Citation() {
                this.cachedSize = -1;
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
                if (this.hasTitleLocalized || !this.titleLocalized.equals("")) {
                    output.writeString(11, this.titleLocalized);
                }
                if (this.hasUrl || !this.url.equals("")) {
                    output.writeString(12, this.url);
                }
                super.writeTo(output);
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final int computeSerializedSize() {
                int size = super.computeSerializedSize();
                if (this.hasTitleLocalized || !this.titleLocalized.equals("")) {
                    size += CodedOutputByteBufferNano.computeStringSize(11, this.titleLocalized);
                }
                if (this.hasUrl || !this.url.equals("")) {
                    return size + CodedOutputByteBufferNano.computeStringSize(12, this.url);
                }
                return size;
            }
        }

        public static Image[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new Image[0];
                    }
                }
            }
            return _emptyArray;
        }

        public Image() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.imageType != 0 || this.hasImageType) {
                output.writeInt32(1, this.imageType);
            }
            if (this.dimension != null) {
                output.writeGroup(2, this.dimension);
            }
            if (this.hasImageUrl || !this.imageUrl.equals("")) {
                output.writeString(5, this.imageUrl);
            }
            if (this.hasAltTextLocalized || !this.altTextLocalized.equals("")) {
                output.writeString(6, this.altTextLocalized);
            }
            if (this.hasSecureUrl || !this.secureUrl.equals("")) {
                output.writeString(7, this.secureUrl);
            }
            if (this.hasPositionInSequence || this.positionInSequence != 0) {
                output.writeInt32(8, this.positionInSequence);
            }
            if (this.hasSupportsFifeUrlOptions || this.supportsFifeUrlOptions) {
                output.writeBool(9, this.supportsFifeUrlOptions);
            }
            if (this.citation != null) {
                output.writeGroup(10, this.citation);
            }
            if (this.hasDurationSeconds || this.durationSeconds != 0) {
                output.writeInt32(14, this.durationSeconds);
            }
            if (this.hasFillColorRgb || !this.fillColorRgb.equals("")) {
                output.writeString(15, this.fillColorRgb);
            }
            if (this.hasAutogen || this.autogen) {
                output.writeBool(16, this.autogen);
            }
            if (this.attribution != null) {
                output.writeMessage(17, this.attribution);
            }
            if (this.hasBackgroundColorRgb || !this.backgroundColorRgb.equals("")) {
                output.writeString(19, this.backgroundColorRgb);
            }
            if (this.palette != null) {
                output.writeMessage(20, this.palette);
            }
            if (this.deviceClass != 0 || this.hasDeviceClass) {
                output.writeInt32(21, this.deviceClass);
            }
            if (this.hasSupportsFifeMonogramOption || this.supportsFifeMonogramOption) {
                output.writeBool(22, this.supportsFifeMonogramOption);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.imageType != 0 || this.hasImageType) {
                size += CodedOutputByteBufferNano.computeInt32Size(1, this.imageType);
            }
            if (this.dimension != null) {
                size += CodedOutputByteBufferNano.computeGroupSize(2, this.dimension);
            }
            if (this.hasImageUrl || !this.imageUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(5, this.imageUrl);
            }
            if (this.hasAltTextLocalized || !this.altTextLocalized.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(6, this.altTextLocalized);
            }
            if (this.hasSecureUrl || !this.secureUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(7, this.secureUrl);
            }
            if (this.hasPositionInSequence || this.positionInSequence != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(8, this.positionInSequence);
            }
            if (this.hasSupportsFifeUrlOptions || this.supportsFifeUrlOptions) {
                size += CodedOutputByteBufferNano.computeTagSize(9) + 1;
            }
            if (this.citation != null) {
                size += CodedOutputByteBufferNano.computeGroupSize(10, this.citation);
            }
            if (this.hasDurationSeconds || this.durationSeconds != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(14, this.durationSeconds);
            }
            if (this.hasFillColorRgb || !this.fillColorRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(15, this.fillColorRgb);
            }
            if (this.hasAutogen || this.autogen) {
                size += CodedOutputByteBufferNano.computeTagSize(16) + 1;
            }
            if (this.attribution != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(17, this.attribution);
            }
            if (this.hasBackgroundColorRgb || !this.backgroundColorRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(19, this.backgroundColorRgb);
            }
            if (this.palette != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(20, this.palette);
            }
            if (this.deviceClass != 0 || this.hasDeviceClass) {
                size += CodedOutputByteBufferNano.computeInt32Size(21, this.deviceClass);
            }
            if (this.hasSupportsFifeMonogramOption || this.supportsFifeMonogramOption) {
                return size + CodedOutputByteBufferNano.computeTagSize(22) + 1;
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class ImagePalette extends MessageNano {
        public String lightVibrantRgb = "";
        public boolean hasLightVibrantRgb = false;
        public String vibrantRgb = "";
        public boolean hasVibrantRgb = false;
        public String darkVibrantRgb = "";
        public boolean hasDarkVibrantRgb = false;
        public String lightMutedRgb = "";
        public boolean hasLightMutedRgb = false;
        public String mutedRgb = "";
        public boolean hasMutedRgb = false;
        public String darkMutedRgb = "";
        public boolean hasDarkMutedRgb = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.lightVibrantRgb = x0.readString();
                        this.hasLightVibrantRgb = true;
                        break;
                    case 18:
                        this.vibrantRgb = x0.readString();
                        this.hasVibrantRgb = true;
                        break;
                    case 26:
                        this.darkVibrantRgb = x0.readString();
                        this.hasDarkVibrantRgb = true;
                        break;
                    case 34:
                        this.lightMutedRgb = x0.readString();
                        this.hasLightMutedRgb = true;
                        break;
                    case 42:
                        this.mutedRgb = x0.readString();
                        this.hasMutedRgb = true;
                        break;
                    case 50:
                        this.darkMutedRgb = x0.readString();
                        this.hasDarkMutedRgb = true;
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

        public ImagePalette() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasLightVibrantRgb || !this.lightVibrantRgb.equals("")) {
                output.writeString(1, this.lightVibrantRgb);
            }
            if (this.hasVibrantRgb || !this.vibrantRgb.equals("")) {
                output.writeString(2, this.vibrantRgb);
            }
            if (this.hasDarkVibrantRgb || !this.darkVibrantRgb.equals("")) {
                output.writeString(3, this.darkVibrantRgb);
            }
            if (this.hasLightMutedRgb || !this.lightMutedRgb.equals("")) {
                output.writeString(4, this.lightMutedRgb);
            }
            if (this.hasMutedRgb || !this.mutedRgb.equals("")) {
                output.writeString(5, this.mutedRgb);
            }
            if (this.hasDarkMutedRgb || !this.darkMutedRgb.equals("")) {
                output.writeString(6, this.darkMutedRgb);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasLightVibrantRgb || !this.lightVibrantRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.lightVibrantRgb);
            }
            if (this.hasVibrantRgb || !this.vibrantRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.vibrantRgb);
            }
            if (this.hasDarkVibrantRgb || !this.darkVibrantRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(3, this.darkVibrantRgb);
            }
            if (this.hasLightMutedRgb || !this.lightMutedRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(4, this.lightMutedRgb);
            }
            if (this.hasMutedRgb || !this.mutedRgb.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(5, this.mutedRgb);
            }
            if (this.hasDarkMutedRgb || !this.darkMutedRgb.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(6, this.darkMutedRgb);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class RentalTerms extends MessageNano {
        public TimePeriod grantPeriod = null;
        public TimePeriod activatePeriod = null;
        public long grantEndTimeSeconds = 0;
        public boolean hasGrantEndTimeSeconds = false;
        public int dEPRECATEDGrantPeriodSeconds = 0;
        public boolean hasDEPRECATEDGrantPeriodSeconds = false;
        public int dEPRECATEDActivatePeriodSeconds = 0;
        public boolean hasDEPRECATEDActivatePeriodSeconds = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.dEPRECATEDGrantPeriodSeconds = x0.readRawVarint32();
                        this.hasDEPRECATEDGrantPeriodSeconds = true;
                        break;
                    case 16:
                        this.dEPRECATEDActivatePeriodSeconds = x0.readRawVarint32();
                        this.hasDEPRECATEDActivatePeriodSeconds = true;
                        break;
                    case 26:
                        if (this.grantPeriod == null) {
                            this.grantPeriod = new TimePeriod();
                        }
                        x0.readMessage(this.grantPeriod);
                        break;
                    case 34:
                        if (this.activatePeriod == null) {
                            this.activatePeriod = new TimePeriod();
                        }
                        x0.readMessage(this.activatePeriod);
                        break;
                    case 40:
                        this.grantEndTimeSeconds = x0.readRawVarint64();
                        this.hasGrantEndTimeSeconds = true;
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

        public RentalTerms() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasDEPRECATEDGrantPeriodSeconds || this.dEPRECATEDGrantPeriodSeconds != 0) {
                output.writeInt32(1, this.dEPRECATEDGrantPeriodSeconds);
            }
            if (this.hasDEPRECATEDActivatePeriodSeconds || this.dEPRECATEDActivatePeriodSeconds != 0) {
                output.writeInt32(2, this.dEPRECATEDActivatePeriodSeconds);
            }
            if (this.grantPeriod != null) {
                output.writeMessage(3, this.grantPeriod);
            }
            if (this.activatePeriod != null) {
                output.writeMessage(4, this.activatePeriod);
            }
            if (this.hasGrantEndTimeSeconds || this.grantEndTimeSeconds != 0) {
                output.writeInt64(5, this.grantEndTimeSeconds);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasDEPRECATEDGrantPeriodSeconds || this.dEPRECATEDGrantPeriodSeconds != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(1, this.dEPRECATEDGrantPeriodSeconds);
            }
            if (this.hasDEPRECATEDActivatePeriodSeconds || this.dEPRECATEDActivatePeriodSeconds != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(2, this.dEPRECATEDActivatePeriodSeconds);
            }
            if (this.grantPeriod != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.grantPeriod);
            }
            if (this.activatePeriod != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(4, this.activatePeriod);
            }
            if (this.hasGrantEndTimeSeconds || this.grantEndTimeSeconds != 0) {
                return size + CodedOutputByteBufferNano.computeInt64Size(5, this.grantEndTimeSeconds);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class TimePeriod extends MessageNano {
        public int unit = 0;
        public boolean hasUnit = false;
        public int count = 0;
        public boolean hasCount = false;

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
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 5:
                            case 6:
                            case 7:
                                this.unit = readRawVarint32;
                                this.hasUnit = true;
                                continue;
                        }
                    case 16:
                        this.count = x0.readRawVarint32();
                        this.hasCount = true;
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

        public TimePeriod() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.unit != 0 || this.hasUnit) {
                output.writeInt32(1, this.unit);
            }
            if (this.hasCount || this.count != 0) {
                output.writeInt32(2, this.count);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.unit != 0 || this.hasUnit) {
                size += CodedOutputByteBufferNano.computeInt32Size(1, this.unit);
            }
            if (this.hasCount || this.count != 0) {
                return size + CodedOutputByteBufferNano.computeInt32Size(2, this.count);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class MonthAndDay extends MessageNano {
        public int month = 0;
        public boolean hasMonth = false;
        public int day = 0;
        public boolean hasDay = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.month = x0.readRawVarint32();
                        this.hasMonth = true;
                        break;
                    case 16:
                        this.day = x0.readRawVarint32();
                        this.hasDay = true;
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

        public MonthAndDay() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasMonth || this.month != 0) {
                output.writeUInt32(1, this.month);
            }
            if (this.hasDay || this.day != 0) {
                output.writeUInt32(2, this.day);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasMonth || this.month != 0) {
                size += CodedOutputByteBufferNano.computeUInt32Size(1, this.month);
            }
            if (this.hasDay || this.day != 0) {
                return size + CodedOutputByteBufferNano.computeUInt32Size(2, this.day);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class SeasonalSubscriptionInfo extends MessageNano {
        public MonthAndDay periodStart = null;
        public MonthAndDay periodEnd = null;
        public boolean prorated = false;
        public boolean hasProrated = false;
        public Payment postTrialConversionPayment = null;

        /* loaded from: classes.dex */
        public static final class Payment extends MessageNano {
            public long micros = 0;
            public boolean hasMicros = false;
            public String currencyCode = "";
            public boolean hasCurrencyCode = false;
            public String formattedAmount = "";
            public boolean hasFormattedAmount = false;
            public TimePeriod period = null;

            @Override // com.google.protobuf.nano.MessageNano
            public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
                while (true) {
                    int readTag = x0.readTag();
                    switch (readTag) {
                        case 0:
                            break;
                        case 8:
                            this.micros = x0.readRawVarint64();
                            this.hasMicros = true;
                            break;
                        case 18:
                            this.currencyCode = x0.readString();
                            this.hasCurrencyCode = true;
                            break;
                        case 26:
                            this.formattedAmount = x0.readString();
                            this.hasFormattedAmount = true;
                            break;
                        case 34:
                            if (this.period == null) {
                                this.period = new TimePeriod();
                            }
                            x0.readMessage(this.period);
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

            public Payment() {
                this.cachedSize = -1;
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
                if (this.hasMicros || this.micros != 0) {
                    output.writeInt64(1, this.micros);
                }
                if (this.hasCurrencyCode || !this.currencyCode.equals("")) {
                    output.writeString(2, this.currencyCode);
                }
                if (this.hasFormattedAmount || !this.formattedAmount.equals("")) {
                    output.writeString(3, this.formattedAmount);
                }
                if (this.period != null) {
                    output.writeMessage(4, this.period);
                }
                super.writeTo(output);
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final int computeSerializedSize() {
                int size = super.computeSerializedSize();
                if (this.hasMicros || this.micros != 0) {
                    size += CodedOutputByteBufferNano.computeInt64Size(1, this.micros);
                }
                if (this.hasCurrencyCode || !this.currencyCode.equals("")) {
                    size += CodedOutputByteBufferNano.computeStringSize(2, this.currencyCode);
                }
                if (this.hasFormattedAmount || !this.formattedAmount.equals("")) {
                    size += CodedOutputByteBufferNano.computeStringSize(3, this.formattedAmount);
                }
                if (this.period != null) {
                    return size + CodedOutputByteBufferNano.computeMessageSize(4, this.period);
                }
                return size;
            }
        }

        public SeasonalSubscriptionInfo() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.periodStart != null) {
                output.writeMessage(1, this.periodStart);
            }
            if (this.periodEnd != null) {
                output.writeMessage(2, this.periodEnd);
            }
            if (this.hasProrated || this.prorated) {
                output.writeBool(4, this.prorated);
            }
            if (this.postTrialConversionPayment != null) {
                output.writeMessage(5, this.postTrialConversionPayment);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.periodStart != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.periodStart);
            }
            if (this.periodEnd != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.periodEnd);
            }
            if (this.hasProrated || this.prorated) {
                size += CodedOutputByteBufferNano.computeTagSize(4) + 1;
            }
            if (this.postTrialConversionPayment != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(5, this.postTrialConversionPayment);
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
                        if (this.periodStart == null) {
                            this.periodStart = new MonthAndDay();
                        }
                        x0.readMessage(this.periodStart);
                        break;
                    case 18:
                        if (this.periodEnd == null) {
                            this.periodEnd = new MonthAndDay();
                        }
                        x0.readMessage(this.periodEnd);
                        break;
                    case 32:
                        this.prorated = x0.readBool();
                        this.hasProrated = true;
                        break;
                    case 42:
                        if (this.postTrialConversionPayment == null) {
                            this.postTrialConversionPayment = new Payment();
                        }
                        x0.readMessage(this.postTrialConversionPayment);
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
    public static final class SubscriptionTerms extends MessageNano {
        public TimePeriod recurringPeriod = null;
        public TimePeriod trialPeriod = null;
        public boolean enableAppSpecifiedTrialPeriod = false;
        public boolean hasEnableAppSpecifiedTrialPeriod = false;
        public String formattedPriceWithRecurrencePeriod = "";
        public boolean hasFormattedPriceWithRecurrencePeriod = false;
        public SeasonalSubscriptionInfo seasonalSubscriptionInfo = null;
        public Docid[] replaceDocid = Docid.emptyArray();
        public SubscriptionReplacement subscriptionReplacement = null;
        public TimePeriod gracePeriod = null;
        public boolean resignup = false;
        public boolean hasResignup = false;
        public long initialValidUntilTimestampMsec = 0;
        public boolean hasInitialValidUntilTimestampMsec = false;
        public String nextPaymentCurrencyCode = "";
        public boolean hasNextPaymentCurrencyCode = false;
        public long nextPaymentPriceMicros = 0;
        public boolean hasNextPaymentPriceMicros = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        if (this.recurringPeriod == null) {
                            this.recurringPeriod = new TimePeriod();
                        }
                        x0.readMessage(this.recurringPeriod);
                        break;
                    case 18:
                        if (this.trialPeriod == null) {
                            this.trialPeriod = new TimePeriod();
                        }
                        x0.readMessage(this.trialPeriod);
                        break;
                    case 26:
                        this.formattedPriceWithRecurrencePeriod = x0.readString();
                        this.hasFormattedPriceWithRecurrencePeriod = true;
                        break;
                    case 34:
                        if (this.seasonalSubscriptionInfo == null) {
                            this.seasonalSubscriptionInfo = new SeasonalSubscriptionInfo();
                        }
                        x0.readMessage(this.seasonalSubscriptionInfo);
                        break;
                    case 42:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 42);
                        if (this.replaceDocid == null) {
                            length = 0;
                        } else {
                            length = this.replaceDocid.length;
                        }
                        Docid[] docidArr = new Docid[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.replaceDocid, 0, docidArr, 0, length);
                        }
                        while (length < docidArr.length - 1) {
                            docidArr[length] = new Docid();
                            x0.readMessage(docidArr[length]);
                            x0.readTag();
                            length++;
                        }
                        docidArr[length] = new Docid();
                        x0.readMessage(docidArr[length]);
                        this.replaceDocid = docidArr;
                        break;
                    case 50:
                        if (this.gracePeriod == null) {
                            this.gracePeriod = new TimePeriod();
                        }
                        x0.readMessage(this.gracePeriod);
                        break;
                    case 56:
                        this.resignup = x0.readBool();
                        this.hasResignup = true;
                        break;
                    case 64:
                        this.initialValidUntilTimestampMsec = x0.readRawVarint64();
                        this.hasInitialValidUntilTimestampMsec = true;
                        break;
                    case 74:
                        this.nextPaymentCurrencyCode = x0.readString();
                        this.hasNextPaymentCurrencyCode = true;
                        break;
                    case 80:
                        this.nextPaymentPriceMicros = x0.readRawVarint64();
                        this.hasNextPaymentPriceMicros = true;
                        break;
                    case 88:
                        this.enableAppSpecifiedTrialPeriod = x0.readBool();
                        this.hasEnableAppSpecifiedTrialPeriod = true;
                        break;
                    case 98:
                        if (this.subscriptionReplacement == null) {
                            this.subscriptionReplacement = new SubscriptionReplacement();
                        }
                        x0.readMessage(this.subscriptionReplacement);
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

        /* loaded from: classes.dex */
        public static final class SubscriptionReplacement extends MessageNano {
            public Docid newDocid = null;
            public Docid[] oldDocid = Docid.emptyArray();
            public boolean keepNextRecurrenceTime = false;
            public boolean hasKeepNextRecurrenceTime = false;
            public boolean replaceOnFirstRecurrence = false;
            public boolean hasReplaceOnFirstRecurrence = false;

            @Override // com.google.protobuf.nano.MessageNano
            public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
                int length;
                while (true) {
                    int readTag = x0.readTag();
                    switch (readTag) {
                        case 0:
                            break;
                        case 10:
                            if (this.newDocid == null) {
                                this.newDocid = new Docid();
                            }
                            x0.readMessage(this.newDocid);
                            break;
                        case 18:
                            int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 18);
                            if (this.oldDocid == null) {
                                length = 0;
                            } else {
                                length = this.oldDocid.length;
                            }
                            Docid[] docidArr = new Docid[repeatedFieldArrayLength + length];
                            if (length != 0) {
                                System.arraycopy(this.oldDocid, 0, docidArr, 0, length);
                            }
                            while (length < docidArr.length - 1) {
                                docidArr[length] = new Docid();
                                x0.readMessage(docidArr[length]);
                                x0.readTag();
                                length++;
                            }
                            docidArr[length] = new Docid();
                            x0.readMessage(docidArr[length]);
                            this.oldDocid = docidArr;
                            break;
                        case 24:
                            this.keepNextRecurrenceTime = x0.readBool();
                            this.hasKeepNextRecurrenceTime = true;
                            break;
                        case 32:
                            this.replaceOnFirstRecurrence = x0.readBool();
                            this.hasReplaceOnFirstRecurrence = true;
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

            public SubscriptionReplacement() {
                this.cachedSize = -1;
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
                if (this.newDocid != null) {
                    output.writeMessage(1, this.newDocid);
                }
                if (this.oldDocid != null && this.oldDocid.length > 0) {
                    for (int i = 0; i < this.oldDocid.length; i++) {
                        Docid element = this.oldDocid[i];
                        if (element != null) {
                            output.writeMessage(2, element);
                        }
                    }
                }
                if (this.hasKeepNextRecurrenceTime || this.keepNextRecurrenceTime) {
                    output.writeBool(3, this.keepNextRecurrenceTime);
                }
                if (this.hasReplaceOnFirstRecurrence || this.replaceOnFirstRecurrence) {
                    output.writeBool(4, this.replaceOnFirstRecurrence);
                }
                super.writeTo(output);
            }

            @Override // com.google.protobuf.nano.MessageNano
            public final int computeSerializedSize() {
                int size = super.computeSerializedSize();
                if (this.newDocid != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(1, this.newDocid);
                }
                if (this.oldDocid != null && this.oldDocid.length > 0) {
                    for (int i = 0; i < this.oldDocid.length; i++) {
                        Docid element = this.oldDocid[i];
                        if (element != null) {
                            size += CodedOutputByteBufferNano.computeMessageSize(2, element);
                        }
                    }
                }
                if (this.hasKeepNextRecurrenceTime || this.keepNextRecurrenceTime) {
                    size += CodedOutputByteBufferNano.computeTagSize(3) + 1;
                }
                if (this.hasReplaceOnFirstRecurrence || this.replaceOnFirstRecurrence) {
                    return size + CodedOutputByteBufferNano.computeTagSize(4) + 1;
                }
                return size;
            }
        }

        public SubscriptionTerms() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.recurringPeriod != null) {
                output.writeMessage(1, this.recurringPeriod);
            }
            if (this.trialPeriod != null) {
                output.writeMessage(2, this.trialPeriod);
            }
            if (this.hasFormattedPriceWithRecurrencePeriod || !this.formattedPriceWithRecurrencePeriod.equals("")) {
                output.writeString(3, this.formattedPriceWithRecurrencePeriod);
            }
            if (this.seasonalSubscriptionInfo != null) {
                output.writeMessage(4, this.seasonalSubscriptionInfo);
            }
            if (this.replaceDocid != null && this.replaceDocid.length > 0) {
                for (int i = 0; i < this.replaceDocid.length; i++) {
                    Docid element = this.replaceDocid[i];
                    if (element != null) {
                        output.writeMessage(5, element);
                    }
                }
            }
            if (this.gracePeriod != null) {
                output.writeMessage(6, this.gracePeriod);
            }
            if (this.hasResignup || this.resignup) {
                output.writeBool(7, this.resignup);
            }
            if (this.hasInitialValidUntilTimestampMsec || this.initialValidUntilTimestampMsec != 0) {
                output.writeInt64(8, this.initialValidUntilTimestampMsec);
            }
            if (this.hasNextPaymentCurrencyCode || !this.nextPaymentCurrencyCode.equals("")) {
                output.writeString(9, this.nextPaymentCurrencyCode);
            }
            if (this.hasNextPaymentPriceMicros || this.nextPaymentPriceMicros != 0) {
                output.writeInt64(10, this.nextPaymentPriceMicros);
            }
            if (this.hasEnableAppSpecifiedTrialPeriod || this.enableAppSpecifiedTrialPeriod) {
                output.writeBool(11, this.enableAppSpecifiedTrialPeriod);
            }
            if (this.subscriptionReplacement != null) {
                output.writeMessage(12, this.subscriptionReplacement);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.recurringPeriod != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.recurringPeriod);
            }
            if (this.trialPeriod != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.trialPeriod);
            }
            if (this.hasFormattedPriceWithRecurrencePeriod || !this.formattedPriceWithRecurrencePeriod.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(3, this.formattedPriceWithRecurrencePeriod);
            }
            if (this.seasonalSubscriptionInfo != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(4, this.seasonalSubscriptionInfo);
            }
            if (this.replaceDocid != null && this.replaceDocid.length > 0) {
                for (int i = 0; i < this.replaceDocid.length; i++) {
                    Docid element = this.replaceDocid[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(5, element);
                    }
                }
            }
            if (this.gracePeriod != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(6, this.gracePeriod);
            }
            if (this.hasResignup || this.resignup) {
                size += CodedOutputByteBufferNano.computeTagSize(7) + 1;
            }
            if (this.hasInitialValidUntilTimestampMsec || this.initialValidUntilTimestampMsec != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(8, this.initialValidUntilTimestampMsec);
            }
            if (this.hasNextPaymentCurrencyCode || !this.nextPaymentCurrencyCode.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(9, this.nextPaymentCurrencyCode);
            }
            if (this.hasNextPaymentPriceMicros || this.nextPaymentPriceMicros != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(10, this.nextPaymentPriceMicros);
            }
            if (this.hasEnableAppSpecifiedTrialPeriod || this.enableAppSpecifiedTrialPeriod) {
                size += CodedOutputByteBufferNano.computeTagSize(11) + 1;
            }
            if (this.subscriptionReplacement != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(12, this.subscriptionReplacement);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class SubscriptionContentTerms extends MessageNano {
        public Docid requiredSubscription = null;

        public SubscriptionContentTerms() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.requiredSubscription != null) {
                output.writeMessage(1, this.requiredSubscription);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.requiredSubscription != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(1, this.requiredSubscription);
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
                        if (this.requiredSubscription == null) {
                            this.requiredSubscription = new Docid();
                        }
                        x0.readMessage(this.requiredSubscription);
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
    public static final class VoucherOfferTerms extends MessageNano {
        public Docid[] voucherDocid = Docid.emptyArray();
        public long voucherPriceMicros = 0;
        public boolean hasVoucherPriceMicros = false;
        public String voucherFormattedAmount = "";
        public boolean hasVoucherFormattedAmount = false;

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
                        if (this.voucherDocid == null) {
                            length = 0;
                        } else {
                            length = this.voucherDocid.length;
                        }
                        Docid[] docidArr = new Docid[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.voucherDocid, 0, docidArr, 0, length);
                        }
                        while (length < docidArr.length - 1) {
                            docidArr[length] = new Docid();
                            x0.readMessage(docidArr[length]);
                            x0.readTag();
                            length++;
                        }
                        docidArr[length] = new Docid();
                        x0.readMessage(docidArr[length]);
                        this.voucherDocid = docidArr;
                        break;
                    case 16:
                        this.voucherPriceMicros = x0.readRawVarint64();
                        this.hasVoucherPriceMicros = true;
                        break;
                    case 26:
                        this.voucherFormattedAmount = x0.readString();
                        this.hasVoucherFormattedAmount = true;
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

        public VoucherOfferTerms() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.voucherDocid != null && this.voucherDocid.length > 0) {
                for (int i = 0; i < this.voucherDocid.length; i++) {
                    Docid element = this.voucherDocid[i];
                    if (element != null) {
                        output.writeMessage(1, element);
                    }
                }
            }
            if (this.hasVoucherPriceMicros || this.voucherPriceMicros != 0) {
                output.writeInt64(2, this.voucherPriceMicros);
            }
            if (this.hasVoucherFormattedAmount || !this.voucherFormattedAmount.equals("")) {
                output.writeString(3, this.voucherFormattedAmount);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.voucherDocid != null && this.voucherDocid.length > 0) {
                for (int i = 0; i < this.voucherDocid.length; i++) {
                    Docid element = this.voucherDocid[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(1, element);
                    }
                }
            }
            if (this.hasVoucherPriceMicros || this.voucherPriceMicros != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(2, this.voucherPriceMicros);
            }
            if (this.hasVoucherFormattedAmount || !this.voucherFormattedAmount.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(3, this.voucherFormattedAmount);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class FamilyShareability extends MessageNano {
        public int state = 0;
        public boolean hasState = false;

        public FamilyShareability() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.state != 0 || this.hasState) {
                output.writeInt32(1, this.state);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.state != 0 || this.hasState) {
                return size + CodedOutputByteBufferNano.computeInt32Size(1, this.state);
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
                    case 8:
                        int readRawVarint32 = x0.readRawVarint32();
                        switch (readRawVarint32) {
                            case 0:
                            case 1:
                            case 2:
                                this.state = readRawVarint32;
                                this.hasState = true;
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
    public static final class GroupLicenseKey extends MessageNano {
        public long dasherCustomerId = 0;
        public boolean hasDasherCustomerId = false;
        public Docid docid = null;
        public int licensedOfferType = 1;
        public boolean hasLicensedOfferType = false;
        public int type = 0;
        public boolean hasType = false;
        public int rentalPeriodDays = 0;
        public boolean hasRentalPeriodDays = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 9:
                        this.dasherCustomerId = x0.readRawLittleEndian64();
                        this.hasDasherCustomerId = true;
                        break;
                    case 18:
                        if (this.docid == null) {
                            this.docid = new Docid();
                        }
                        x0.readMessage(this.docid);
                        break;
                    case 24:
                        int readRawVarint32 = x0.readRawVarint32();
                        switch (readRawVarint32) {
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 6:
                            case 7:
                            case 8:
                            case 9:
                            case 10:
                            case 11:
                            case 12:
                            case 13:
                                this.licensedOfferType = readRawVarint32;
                                this.hasLicensedOfferType = true;
                                continue;
                        }
                    case 32:
                        int readRawVarint322 = x0.readRawVarint32();
                        switch (readRawVarint322) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                                this.type = readRawVarint322;
                                this.hasType = true;
                                continue;
                        }
                    case 40:
                        this.rentalPeriodDays = x0.readRawVarint32();
                        this.hasRentalPeriodDays = true;
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

        public GroupLicenseKey() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasDasherCustomerId || this.dasherCustomerId != 0) {
                output.writeFixed64(1, this.dasherCustomerId);
            }
            if (this.docid != null) {
                output.writeMessage(2, this.docid);
            }
            if (this.licensedOfferType != 1 || this.hasLicensedOfferType) {
                output.writeInt32(3, this.licensedOfferType);
            }
            if (this.type != 0 || this.hasType) {
                output.writeInt32(4, this.type);
            }
            if (this.hasRentalPeriodDays || this.rentalPeriodDays != 0) {
                output.writeInt32(5, this.rentalPeriodDays);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasDasherCustomerId || this.dasherCustomerId != 0) {
                size += CodedOutputByteBufferNano.computeTagSize(1) + 8;
            }
            if (this.docid != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.docid);
            }
            if (this.licensedOfferType != 1 || this.hasLicensedOfferType) {
                size += CodedOutputByteBufferNano.computeInt32Size(3, this.licensedOfferType);
            }
            if (this.type != 0 || this.hasType) {
                size += CodedOutputByteBufferNano.computeInt32Size(4, this.type);
            }
            if (this.hasRentalPeriodDays || this.rentalPeriodDays != 0) {
                return size + CodedOutputByteBufferNano.computeInt32Size(5, this.rentalPeriodDays);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class LicensedDocumentInfo extends MessageNano {
        public long[] gaiaGroupId = WireFormatNano.EMPTY_LONG_ARRAY;
        public String groupLicenseCheckoutOrderId = "";
        public boolean hasGroupLicenseCheckoutOrderId = false;
        public GroupLicenseKey groupLicenseKey = null;
        public long assignedByGaiaId = 0;
        public boolean hasAssignedByGaiaId = false;
        public String dEPRECATEDAssignmentId = "";
        public boolean hasDEPRECATEDAssignmentId = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 9:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 9);
                        if (this.gaiaGroupId == null) {
                            length = 0;
                        } else {
                            length = this.gaiaGroupId.length;
                        }
                        long[] jArr = new long[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.gaiaGroupId, 0, jArr, 0, length);
                        }
                        while (length < jArr.length - 1) {
                            jArr[length] = x0.readRawLittleEndian64();
                            x0.readTag();
                            length++;
                        }
                        jArr[length] = x0.readRawLittleEndian64();
                        this.gaiaGroupId = jArr;
                        break;
                    case 10:
                        int readRawVarint32 = x0.readRawVarint32();
                        int pushLimit = x0.pushLimit(readRawVarint32);
                        int i = readRawVarint32 / 8;
                        int length2 = this.gaiaGroupId == null ? 0 : this.gaiaGroupId.length;
                        long[] jArr2 = new long[i + length2];
                        if (length2 != 0) {
                            System.arraycopy(this.gaiaGroupId, 0, jArr2, 0, length2);
                        }
                        while (length2 < jArr2.length) {
                            jArr2[length2] = x0.readRawLittleEndian64();
                            length2++;
                        }
                        this.gaiaGroupId = jArr2;
                        x0.popLimit(pushLimit);
                        break;
                    case 18:
                        this.groupLicenseCheckoutOrderId = x0.readString();
                        this.hasGroupLicenseCheckoutOrderId = true;
                        break;
                    case 26:
                        if (this.groupLicenseKey == null) {
                            this.groupLicenseKey = new GroupLicenseKey();
                        }
                        x0.readMessage(this.groupLicenseKey);
                        break;
                    case 33:
                        this.assignedByGaiaId = x0.readRawLittleEndian64();
                        this.hasAssignedByGaiaId = true;
                        break;
                    case 42:
                        this.dEPRECATEDAssignmentId = x0.readString();
                        this.hasDEPRECATEDAssignmentId = true;
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

        public LicensedDocumentInfo() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.gaiaGroupId != null && this.gaiaGroupId.length > 0) {
                for (int i = 0; i < this.gaiaGroupId.length; i++) {
                    output.writeFixed64(1, this.gaiaGroupId[i]);
                }
            }
            if (this.hasGroupLicenseCheckoutOrderId || !this.groupLicenseCheckoutOrderId.equals("")) {
                output.writeString(2, this.groupLicenseCheckoutOrderId);
            }
            if (this.groupLicenseKey != null) {
                output.writeMessage(3, this.groupLicenseKey);
            }
            if (this.hasAssignedByGaiaId || this.assignedByGaiaId != 0) {
                output.writeFixed64(4, this.assignedByGaiaId);
            }
            if (this.hasDEPRECATEDAssignmentId || !this.dEPRECATEDAssignmentId.equals("")) {
                output.writeString(5, this.dEPRECATEDAssignmentId);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.gaiaGroupId != null && this.gaiaGroupId.length > 0) {
                size = size + (this.gaiaGroupId.length * 8) + (this.gaiaGroupId.length * 1);
            }
            if (this.hasGroupLicenseCheckoutOrderId || !this.groupLicenseCheckoutOrderId.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.groupLicenseCheckoutOrderId);
            }
            if (this.groupLicenseKey != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.groupLicenseKey);
            }
            if (this.hasAssignedByGaiaId || this.assignedByGaiaId != 0) {
                size += CodedOutputByteBufferNano.computeTagSize(4) + 8;
            }
            if (this.hasDEPRECATEDAssignmentId || !this.dEPRECATEDAssignmentId.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(5, this.dEPRECATEDAssignmentId);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class LicenseTerms extends MessageNano {
        public GroupLicenseKey groupLicenseKey = null;

        public LicenseTerms() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.groupLicenseKey != null) {
                output.writeMessage(1, this.groupLicenseKey);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.groupLicenseKey != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(1, this.groupLicenseKey);
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
                        if (this.groupLicenseKey == null) {
                            this.groupLicenseKey = new GroupLicenseKey();
                        }
                        x0.readMessage(this.groupLicenseKey);
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
    public static final class RedemptionRecordKey extends MessageNano {
        public int type = 1;
        public boolean hasType = false;
        public long publisherId = 0;
        public boolean hasPublisherId = false;
        public long campaignId = 0;
        public boolean hasCampaignId = false;
        public long codeGroupId = 0;
        public boolean hasCodeGroupId = false;
        public long recordId = 0;
        public boolean hasRecordId = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 8:
                        this.publisherId = x0.readRawVarint64();
                        this.hasPublisherId = true;
                        break;
                    case 16:
                        this.campaignId = x0.readRawVarint64();
                        this.hasCampaignId = true;
                        break;
                    case 24:
                        this.codeGroupId = x0.readRawVarint64();
                        this.hasCodeGroupId = true;
                        break;
                    case 32:
                        this.recordId = x0.readRawVarint64();
                        this.hasRecordId = true;
                        break;
                    case 40:
                        int readRawVarint32 = x0.readRawVarint32();
                        switch (readRawVarint32) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                                this.type = readRawVarint32;
                                this.hasType = true;
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

        public RedemptionRecordKey() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasPublisherId || this.publisherId != 0) {
                output.writeInt64(1, this.publisherId);
            }
            if (this.hasCampaignId || this.campaignId != 0) {
                output.writeInt64(2, this.campaignId);
            }
            if (this.hasCodeGroupId || this.codeGroupId != 0) {
                output.writeInt64(3, this.codeGroupId);
            }
            if (this.hasRecordId || this.recordId != 0) {
                output.writeInt64(4, this.recordId);
            }
            if (this.type != 1 || this.hasType) {
                output.writeInt32(5, this.type);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasPublisherId || this.publisherId != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(1, this.publisherId);
            }
            if (this.hasCampaignId || this.campaignId != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(2, this.campaignId);
            }
            if (this.hasCodeGroupId || this.codeGroupId != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(3, this.codeGroupId);
            }
            if (this.hasRecordId || this.recordId != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(4, this.recordId);
            }
            if (this.type != 1 || this.hasType) {
                return size + CodedOutputByteBufferNano.computeInt32Size(5, this.type);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class VoucherId extends MessageNano {
        public Docid voucherDocid = null;
        public RedemptionRecordKey key = null;

        public VoucherId() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.voucherDocid != null) {
                output.writeMessage(1, this.voucherDocid);
            }
            if (this.key != null) {
                output.writeMessage(2, this.key);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.voucherDocid != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.voucherDocid);
            }
            if (this.key != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(2, this.key);
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
                        if (this.voucherDocid == null) {
                            this.voucherDocid = new Docid();
                        }
                        x0.readMessage(this.voucherDocid);
                        break;
                    case 18:
                        if (this.key == null) {
                            this.key = new RedemptionRecordKey();
                        }
                        x0.readMessage(this.key);
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
    public static final class Install extends MessageNano {
        private static volatile Install[] _emptyArray;
        public long androidId = 0;
        public boolean hasAndroidId = false;
        public int version = 0;
        public boolean hasVersion = false;
        public boolean bundled = false;
        public boolean hasBundled = false;
        public boolean pending = false;
        public boolean hasPending = false;
        public long lastUpdateTimestampMillis = 0;
        public boolean hasLastUpdateTimestampMillis = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 9:
                        this.androidId = x0.readRawLittleEndian64();
                        this.hasAndroidId = true;
                        break;
                    case 16:
                        this.version = x0.readRawVarint32();
                        this.hasVersion = true;
                        break;
                    case 24:
                        this.bundled = x0.readBool();
                        this.hasBundled = true;
                        break;
                    case 32:
                        this.pending = x0.readBool();
                        this.hasPending = true;
                        break;
                    case 40:
                        this.lastUpdateTimestampMillis = x0.readRawVarint64();
                        this.hasLastUpdateTimestampMillis = true;
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

        public static Install[] emptyArray() {
            if (_emptyArray == null) {
                synchronized (InternalNano.LAZY_INIT_LOCK) {
                    if (_emptyArray == null) {
                        _emptyArray = new Install[0];
                    }
                }
            }
            return _emptyArray;
        }

        public Install() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasAndroidId || this.androidId != 0) {
                output.writeFixed64(1, this.androidId);
            }
            if (this.hasVersion || this.version != 0) {
                output.writeInt32(2, this.version);
            }
            if (this.hasBundled || this.bundled) {
                output.writeBool(3, this.bundled);
            }
            if (this.hasPending || this.pending) {
                output.writeBool(4, this.pending);
            }
            if (this.hasLastUpdateTimestampMillis || this.lastUpdateTimestampMillis != 0) {
                output.writeInt64(5, this.lastUpdateTimestampMillis);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasAndroidId || this.androidId != 0) {
                size += CodedOutputByteBufferNano.computeTagSize(1) + 8;
            }
            if (this.hasVersion || this.version != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(2, this.version);
            }
            if (this.hasBundled || this.bundled) {
                size += CodedOutputByteBufferNano.computeTagSize(3) + 1;
            }
            if (this.hasPending || this.pending) {
                size += CodedOutputByteBufferNano.computeTagSize(4) + 1;
            }
            if (this.hasLastUpdateTimestampMillis || this.lastUpdateTimestampMillis != 0) {
                return size + CodedOutputByteBufferNano.computeInt64Size(5, this.lastUpdateTimestampMillis);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class SignedData extends MessageNano {
        public String signedData = "";
        public boolean hasSignedData = false;
        public String signature = "";
        public boolean hasSignature = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.signedData = x0.readString();
                        this.hasSignedData = true;
                        break;
                    case 18:
                        this.signature = x0.readString();
                        this.hasSignature = true;
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

        public SignedData() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasSignedData || !this.signedData.equals("")) {
                output.writeString(1, this.signedData);
            }
            if (this.hasSignature || !this.signature.equals("")) {
                output.writeString(2, this.signature);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasSignedData || !this.signedData.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.signedData);
            }
            if (this.hasSignature || !this.signature.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(2, this.signature);
            }
            return size;
        }
    }

    /* loaded from: classes.dex */
    public static final class Attribution extends MessageNano {
        public String sourceTitle = "";
        public boolean hasSourceTitle = false;
        public String sourceUrl = "";
        public boolean hasSourceUrl = false;
        public String licenseTitle = "";
        public boolean hasLicenseTitle = false;
        public String licenseUrl = "";
        public boolean hasLicenseUrl = false;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        this.sourceTitle = x0.readString();
                        this.hasSourceTitle = true;
                        break;
                    case 18:
                        this.sourceUrl = x0.readString();
                        this.hasSourceUrl = true;
                        break;
                    case 26:
                        this.licenseTitle = x0.readString();
                        this.hasLicenseTitle = true;
                        break;
                    case 34:
                        this.licenseUrl = x0.readString();
                        this.hasLicenseUrl = true;
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

        public Attribution() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.hasSourceTitle || !this.sourceTitle.equals("")) {
                output.writeString(1, this.sourceTitle);
            }
            if (this.hasSourceUrl || !this.sourceUrl.equals("")) {
                output.writeString(2, this.sourceUrl);
            }
            if (this.hasLicenseTitle || !this.licenseTitle.equals("")) {
                output.writeString(3, this.licenseTitle);
            }
            if (this.hasLicenseUrl || !this.licenseUrl.equals("")) {
                output.writeString(4, this.licenseUrl);
            }
            super.writeTo(output);
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.hasSourceTitle || !this.sourceTitle.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(1, this.sourceTitle);
            }
            if (this.hasSourceUrl || !this.sourceUrl.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(2, this.sourceUrl);
            }
            if (this.hasLicenseTitle || !this.licenseTitle.equals("")) {
                size += CodedOutputByteBufferNano.computeStringSize(3, this.licenseTitle);
            }
            if (this.hasLicenseUrl || !this.licenseUrl.equals("")) {
                return size + CodedOutputByteBufferNano.computeStringSize(4, this.licenseUrl);
            }
            return size;
        }
    }
}
