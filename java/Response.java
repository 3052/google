package com.google.android.finsky.protos;

import com.google.android.finsky.protos.Browse;
import com.google.android.finsky.protos.Buy;
import com.google.android.finsky.protos.BuyInstruments;
import com.google.android.finsky.protos.CarrierBilling;
import com.google.android.finsky.protos.ContentFilters;
import com.google.android.finsky.protos.ContentFlagging;
import com.google.android.finsky.protos.Details;
import com.google.android.finsky.protos.Log;
import com.google.android.finsky.protos.ModifyLibrary;
import com.google.android.finsky.protos.Preloads;
import com.google.android.finsky.protos.PromoCode;
import com.google.android.finsky.protos.Purchase;
import com.google.android.finsky.protos.ResolveLink;
import com.google.android.finsky.protos.Restore;
import com.google.android.finsky.protos.Search;
import com.google.android.finsky.protos.SearchSuggest;
import com.google.android.finsky.protos.SharingSetting;
import com.google.android.finsky.protos.Toc;
import com.google.android.finsky.protos.UploadDeviceConfig;
import com.google.android.finsky.protos.UserActivity;
import com.google.android.finsky.protos.UserRefund;
import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.InvalidProtocolBufferNanoException;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public interface Response {

    /* loaded from: classes.dex */
    public static final class ResponseWrapper extends MessageNano {
        public Payload payload = null;
        public ServerCommands commands = null;
        public PreFetch[] preFetch = PreFetch.emptyArray();
        public Notification[] notification = Notification.emptyArray();
        public ServerMetadata serverMetadata = null;
        public Targets targets = null;
        public ServerCookies serverCookies = null;

        @Override // com.google.protobuf.nano.MessageNano
        public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
            int length;
            int length2;
            while (true) {
                int readTag = x0.readTag();
                switch (readTag) {
                    case 0:
                        break;
                    case 10:
                        if (this.payload == null) {
                            this.payload = new Payload();
                        }
                        x0.readMessage(this.payload);
                        break;
                    case 18:
                        if (this.commands == null) {
                            this.commands = new ServerCommands();
                        }
                        x0.readMessage(this.commands);
                        break;
                    case 26:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 26);
                        if (this.preFetch == null) {
                            length2 = 0;
                        } else {
                            length2 = this.preFetch.length;
                        }
                        PreFetch[] preFetchArr = new PreFetch[repeatedFieldArrayLength + length2];
                        if (length2 != 0) {
                            System.arraycopy(this.preFetch, 0, preFetchArr, 0, length2);
                        }
                        while (length2 < preFetchArr.length - 1) {
                            preFetchArr[length2] = new PreFetch();
                            x0.readMessage(preFetchArr[length2]);
                            x0.readTag();
                            length2++;
                        }
                        preFetchArr[length2] = new PreFetch();
                        x0.readMessage(preFetchArr[length2]);
                        this.preFetch = preFetchArr;
                        break;
                    case 34:
                        int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 34);
                        if (this.notification == null) {
                            length = 0;
                        } else {
                            length = this.notification.length;
                        }
                        Notification[] notificationArr = new Notification[repeatedFieldArrayLength2 + length];
                        if (length != 0) {
                            System.arraycopy(this.notification, 0, notificationArr, 0, length);
                        }
                        while (length < notificationArr.length - 1) {
                            notificationArr[length] = new Notification();
                            x0.readMessage(notificationArr[length]);
                            x0.readTag();
                            length++;
                        }
                        notificationArr[length] = new Notification();
                        x0.readMessage(notificationArr[length]);
                        this.notification = notificationArr;
                        break;
                    case 42:
                        if (this.serverMetadata == null) {
                            this.serverMetadata = new ServerMetadata();
                        }
                        x0.readMessage(this.serverMetadata);
                        break;
                    case 50:
                        if (this.targets == null) {
                            this.targets = new Targets();
                        }
                        x0.readMessage(this.targets);
                        break;
                    case 58:
                        if (this.serverCookies == null) {
                            this.serverCookies = new ServerCookies();
                        }
                        x0.readMessage(this.serverCookies);
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

        public ResponseWrapper() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.payload != null) {
                output.writeMessage(1, this.payload);
            }
            if (this.commands != null) {
                output.writeMessage(2, this.commands);
            }
            if (this.preFetch != null && this.preFetch.length > 0) {
                for (int i = 0; i < this.preFetch.length; i++) {
                    PreFetch element = this.preFetch[i];
                    if (element != null) {
                        output.writeMessage(3, element);
                    }
                }
            }
            if (this.notification != null && this.notification.length > 0) {
                for (int i2 = 0; i2 < this.notification.length; i2++) {
                    Notification element2 = this.notification[i2];
                    if (element2 != null) {
                        output.writeMessage(4, element2);
                    }
                }
            }
            if (this.serverMetadata != null) {
                output.writeMessage(5, this.serverMetadata);
            }
            if (this.targets != null) {
                output.writeMessage(6, this.targets);
            }
            if (this.serverCookies != null) {
                output.writeMessage(7, this.serverCookies);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.payload != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.payload);
            }
            if (this.commands != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.commands);
            }
            if (this.preFetch != null && this.preFetch.length > 0) {
                for (int i = 0; i < this.preFetch.length; i++) {
                    PreFetch element = this.preFetch[i];
                    if (element != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(3, element);
                    }
                }
            }
            if (this.notification != null && this.notification.length > 0) {
                for (int i2 = 0; i2 < this.notification.length; i2++) {
                    Notification element2 = this.notification[i2];
                    if (element2 != null) {
                        size += CodedOutputByteBufferNano.computeMessageSize(4, element2);
                    }
                }
            }
            if (this.serverMetadata != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(5, this.serverMetadata);
            }
            if (this.targets != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(6, this.targets);
            }
            if (this.serverCookies != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(7, this.serverCookies);
            }
            return size;
        }

        public static ResponseWrapper parseFrom(byte[] data) throws InvalidProtocolBufferNanoException {
            return (ResponseWrapper) MessageNano.mergeFrom$1ec43da(new ResponseWrapper(), data, data.length);
        }
    }

    /* loaded from: classes.dex */
    public static final class Payload extends MessageNano {
        public ListResponse listResponse = null;
        public Details.DetailsResponse detailsResponse = null;
        public ReviewResponse reviewResponse = null;
        public Buy.BuyResponse buyResponse = null;
        public Search.SearchResponse searchResponse = null;
        public Toc.TocResponse tocResponse = null;
        public Browse.BrowseResponse browseResponse = null;
        public Buy.PurchaseStatusResponse purchaseStatusResponse = null;
        public BuyInstruments.UpdateInstrumentResponse updateInstrumentResponse = null;
        public Log.LogResponse logResponse = null;
        public BuyInstruments.CheckInstrumentResponse checkInstrumentResponse = null;
        public PlusOneResponse plusOneResponse = null;
        public ContentFlagging.FlagContentResponse flagContentResponse = null;
        public AckNotificationResponse ackNotificationResponse = null;
        public CarrierBilling.InitiateAssociationResponse initiateAssociationResponse = null;
        public CarrierBilling.VerifyAssociationResponse verifyAssociationResponse = null;
        public LibraryReplicationResponse libraryReplicationResponse = null;
        public RevokeResponse revokeResponse = null;
        public Details.BulkDetailsResponse bulkDetailsResponse = null;
        public ResolveLink.ResolvedLink resolveLinkResponse = null;
        public DeliveryResponse deliveryResponse = null;
        public AcceptTosResponse acceptTosResponse = null;
        public RateSuggestedContentResponse rateSuggestedContentResponse = null;
        public CheckPromoOfferResponse checkPromoOfferResponse = null;
        public BuyInstruments.InstrumentSetupInfoResponse instrumentSetupInfoResponse = null;
        public BuyInstruments.RedeemGiftCardResponse redeemGiftCardResponse = null;
        public ModifyLibrary.ModifyLibraryResponse modifyLibraryResponse = null;
        public UploadDeviceConfig.UploadDeviceConfigResponse uploadDeviceConfigResponse = null;
        public PlusProfileResponse plusProfileResponse = null;
        public ConsumePurchaseResponse consumePurchaseResponse = null;
        public BuyInstruments.BillingProfileResponse billingProfileResponse = null;
        public Purchase.PreparePurchaseResponse preparePurchaseResponse = null;
        public Purchase.CommitPurchaseResponse commitPurchaseResponse = null;
        public DebugSettingsResponse debugSettingsResponse = null;
        public BuyInstruments.CheckIabPromoResponse checkIabPromoResponse = null;
        public UserActivity.UserActivitySettingsResponse userActivitySettingsResponse = null;
        public UserActivity.RecordUserActivityResponse recordUserActivityResponse = null;
        public PromoCode.RedeemCodeResponse redeemCodeResponse = null;
        public SelfUpdateResponse selfUpdateResponse = null;
        public SearchSuggest.SearchSuggestResponse searchSuggestResponse = null;
        public BuyInstruments.GetInitialInstrumentFlowStateResponse getInitialInstrumentFlowStateResponse = null;
        public BuyInstruments.CreateInstrumentResponse createInstrumentResponse = null;
        public ChallengeResponse challengeResponse = null;
        public Restore.GetBackupDeviceChoicesResponse backupDeviceChoicesResponse = null;
        public Restore.GetBackupDocumentChoicesResponse backupDocumentChoicesResponse = null;
        public EarlyUpdateResponse earlyUpdateResponse = null;
        public Preloads.PreloadsResponse preloadsResponse = null;
        public MyAccountResponse myAccountResponse = null;
        public ContentFilters.ContentFilterSettingsResponse contentFilterResponse = null;
        public ExperimentsResponse experimentsResponse = null;
        public SurveyResponse surveyResponse = null;
        public PingResponse pingResponse = null;
        public UpdateUserSettingResponse updateUserSettingResponse = null;
        public GetUserSettingsResponse getUserSettingsResponse = null;
        public SharingSetting.GetFamilySharingSettingsResponse getSharingSettingsResponse = null;
        public SharingSetting.UpdateFamilySharingSettingsResponse updateSharingSettingsResponse = null;
        public ReviewSnippetsResponse reviewSnippetsResponse = null;
        public DocumentSharingStateResponse documentSharingStateResponse = null;
        public GetFamilyPurchaseSettingResponse getFamilyPurchaseSettingResponse = null;
        public SetFamilyPurchaseSettingResponse setFamilyPurchaseSettingReponse = null;
        public DismissSurveyResponse dismissSurveyResponse = null;
        public Purchase.InstantPurchaseResponse instantPurchaseResponse = null;
        public FamilyFopResponse familyFopResponse = null;
        public MonetaryGiftDetailsResponse monetaryGiftDetailsResponse = null;
        public GiftDialogDetailsResponse giftDialogDetailsResponse = null;
        public Purchase.InAppPurchaseHistoryResponse inAppPurchaseHistoryResponse = null;
        public UserRefund.PrepareUserRefundResponse prepareUserRefundResponse = null;
        public UserRefund.CommitUserRefundResponse commitUserRefundResponse = null;
        public ModuleDeliveryResponse moduleDeliveryResponse = null;

        public Payload() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.listResponse != null) {
                output.writeMessage(1, this.listResponse);
            }
            if (this.detailsResponse != null) {
                output.writeMessage(2, this.detailsResponse);
            }
            if (this.reviewResponse != null) {
                output.writeMessage(3, this.reviewResponse);
            }
            if (this.buyResponse != null) {
                output.writeMessage(4, this.buyResponse);
            }
            if (this.searchResponse != null) {
                output.writeMessage(5, this.searchResponse);
            }
            if (this.tocResponse != null) {
                output.writeMessage(6, this.tocResponse);
            }
            if (this.browseResponse != null) {
                output.writeMessage(7, this.browseResponse);
            }
            if (this.purchaseStatusResponse != null) {
                output.writeMessage(8, this.purchaseStatusResponse);
            }
            if (this.updateInstrumentResponse != null) {
                output.writeMessage(9, this.updateInstrumentResponse);
            }
            if (this.logResponse != null) {
                output.writeMessage(10, this.logResponse);
            }
            if (this.checkInstrumentResponse != null) {
                output.writeMessage(11, this.checkInstrumentResponse);
            }
            if (this.plusOneResponse != null) {
                output.writeMessage(12, this.plusOneResponse);
            }
            if (this.flagContentResponse != null) {
                output.writeMessage(13, this.flagContentResponse);
            }
            if (this.ackNotificationResponse != null) {
                output.writeMessage(14, this.ackNotificationResponse);
            }
            if (this.initiateAssociationResponse != null) {
                output.writeMessage(15, this.initiateAssociationResponse);
            }
            if (this.verifyAssociationResponse != null) {
                output.writeMessage(16, this.verifyAssociationResponse);
            }
            if (this.libraryReplicationResponse != null) {
                output.writeMessage(17, this.libraryReplicationResponse);
            }
            if (this.revokeResponse != null) {
                output.writeMessage(18, this.revokeResponse);
            }
            if (this.bulkDetailsResponse != null) {
                output.writeMessage(19, this.bulkDetailsResponse);
            }
            if (this.resolveLinkResponse != null) {
                output.writeMessage(20, this.resolveLinkResponse);
            }
            if (this.deliveryResponse != null) {
                output.writeMessage(21, this.deliveryResponse);
            }
            if (this.acceptTosResponse != null) {
                output.writeMessage(22, this.acceptTosResponse);
            }
            if (this.rateSuggestedContentResponse != null) {
                output.writeMessage(23, this.rateSuggestedContentResponse);
            }
            if (this.checkPromoOfferResponse != null) {
                output.writeMessage(24, this.checkPromoOfferResponse);
            }
            if (this.instrumentSetupInfoResponse != null) {
                output.writeMessage(25, this.instrumentSetupInfoResponse);
            }
            if (this.redeemGiftCardResponse != null) {
                output.writeMessage(26, this.redeemGiftCardResponse);
            }
            if (this.modifyLibraryResponse != null) {
                output.writeMessage(27, this.modifyLibraryResponse);
            }
            if (this.uploadDeviceConfigResponse != null) {
                output.writeMessage(28, this.uploadDeviceConfigResponse);
            }
            if (this.plusProfileResponse != null) {
                output.writeMessage(29, this.plusProfileResponse);
            }
            if (this.consumePurchaseResponse != null) {
                output.writeMessage(30, this.consumePurchaseResponse);
            }
            if (this.billingProfileResponse != null) {
                output.writeMessage(31, this.billingProfileResponse);
            }
            if (this.preparePurchaseResponse != null) {
                output.writeMessage(32, this.preparePurchaseResponse);
            }
            if (this.commitPurchaseResponse != null) {
                output.writeMessage(33, this.commitPurchaseResponse);
            }
            if (this.debugSettingsResponse != null) {
                output.writeMessage(34, this.debugSettingsResponse);
            }
            if (this.checkIabPromoResponse != null) {
                output.writeMessage(35, this.checkIabPromoResponse);
            }
            if (this.userActivitySettingsResponse != null) {
                output.writeMessage(36, this.userActivitySettingsResponse);
            }
            if (this.recordUserActivityResponse != null) {
                output.writeMessage(37, this.recordUserActivityResponse);
            }
            if (this.redeemCodeResponse != null) {
                output.writeMessage(38, this.redeemCodeResponse);
            }
            if (this.selfUpdateResponse != null) {
                output.writeMessage(39, this.selfUpdateResponse);
            }
            if (this.searchSuggestResponse != null) {
                output.writeMessage(40, this.searchSuggestResponse);
            }
            if (this.getInitialInstrumentFlowStateResponse != null) {
                output.writeMessage(41, this.getInitialInstrumentFlowStateResponse);
            }
            if (this.createInstrumentResponse != null) {
                output.writeMessage(42, this.createInstrumentResponse);
            }
            if (this.challengeResponse != null) {
                output.writeMessage(43, this.challengeResponse);
            }
            if (this.backupDeviceChoicesResponse != null) {
                output.writeMessage(44, this.backupDeviceChoicesResponse);
            }
            if (this.backupDocumentChoicesResponse != null) {
                output.writeMessage(45, this.backupDocumentChoicesResponse);
            }
            if (this.earlyUpdateResponse != null) {
                output.writeMessage(46, this.earlyUpdateResponse);
            }
            if (this.preloadsResponse != null) {
                output.writeMessage(47, this.preloadsResponse);
            }
            if (this.myAccountResponse != null) {
                output.writeMessage(48, this.myAccountResponse);
            }
            if (this.contentFilterResponse != null) {
                output.writeMessage(49, this.contentFilterResponse);
            }
            if (this.experimentsResponse != null) {
                output.writeMessage(50, this.experimentsResponse);
            }
            if (this.surveyResponse != null) {
                output.writeMessage(51, this.surveyResponse);
            }
            if (this.pingResponse != null) {
                output.writeMessage(52, this.pingResponse);
            }
            if (this.updateUserSettingResponse != null) {
                output.writeMessage(53, this.updateUserSettingResponse);
            }
            if (this.getUserSettingsResponse != null) {
                output.writeMessage(54, this.getUserSettingsResponse);
            }
            if (this.getSharingSettingsResponse != null) {
                output.writeMessage(56, this.getSharingSettingsResponse);
            }
            if (this.updateSharingSettingsResponse != null) {
                output.writeMessage(57, this.updateSharingSettingsResponse);
            }
            if (this.reviewSnippetsResponse != null) {
                output.writeMessage(58, this.reviewSnippetsResponse);
            }
            if (this.documentSharingStateResponse != null) {
                output.writeMessage(59, this.documentSharingStateResponse);
            }
            if (this.getFamilyPurchaseSettingResponse != null) {
                output.writeMessage(60, this.getFamilyPurchaseSettingResponse);
            }
            if (this.setFamilyPurchaseSettingReponse != null) {
                output.writeMessage(61, this.setFamilyPurchaseSettingReponse);
            }
            if (this.dismissSurveyResponse != null) {
                output.writeMessage(62, this.dismissSurveyResponse);
            }
            if (this.instantPurchaseResponse != null) {
                output.writeMessage(63, this.instantPurchaseResponse);
            }
            if (this.familyFopResponse != null) {
                output.writeMessage(64, this.familyFopResponse);
            }
            if (this.monetaryGiftDetailsResponse != null) {
                output.writeMessage(65, this.monetaryGiftDetailsResponse);
            }
            if (this.giftDialogDetailsResponse != null) {
                output.writeMessage(66, this.giftDialogDetailsResponse);
            }
            if (this.inAppPurchaseHistoryResponse != null) {
                output.writeMessage(67, this.inAppPurchaseHistoryResponse);
            }
            if (this.prepareUserRefundResponse != null) {
                output.writeMessage(68, this.prepareUserRefundResponse);
            }
            if (this.commitUserRefundResponse != null) {
                output.writeMessage(69, this.commitUserRefundResponse);
            }
            if (this.moduleDeliveryResponse != null) {
                output.writeMessage(70, this.moduleDeliveryResponse);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.listResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(1, this.listResponse);
            }
            if (this.detailsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(2, this.detailsResponse);
            }
            if (this.reviewResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(3, this.reviewResponse);
            }
            if (this.buyResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(4, this.buyResponse);
            }
            if (this.searchResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(5, this.searchResponse);
            }
            if (this.tocResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(6, this.tocResponse);
            }
            if (this.browseResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(7, this.browseResponse);
            }
            if (this.purchaseStatusResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(8, this.purchaseStatusResponse);
            }
            if (this.updateInstrumentResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(9, this.updateInstrumentResponse);
            }
            if (this.logResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(10, this.logResponse);
            }
            if (this.checkInstrumentResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(11, this.checkInstrumentResponse);
            }
            if (this.plusOneResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(12, this.plusOneResponse);
            }
            if (this.flagContentResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(13, this.flagContentResponse);
            }
            if (this.ackNotificationResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(14, this.ackNotificationResponse);
            }
            if (this.initiateAssociationResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(15, this.initiateAssociationResponse);
            }
            if (this.verifyAssociationResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(16, this.verifyAssociationResponse);
            }
            if (this.libraryReplicationResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(17, this.libraryReplicationResponse);
            }
            if (this.revokeResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(18, this.revokeResponse);
            }
            if (this.bulkDetailsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(19, this.bulkDetailsResponse);
            }
            if (this.resolveLinkResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(20, this.resolveLinkResponse);
            }
            if (this.deliveryResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(21, this.deliveryResponse);
            }
            if (this.acceptTosResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(22, this.acceptTosResponse);
            }
            if (this.rateSuggestedContentResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(23, this.rateSuggestedContentResponse);
            }
            if (this.checkPromoOfferResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(24, this.checkPromoOfferResponse);
            }
            if (this.instrumentSetupInfoResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(25, this.instrumentSetupInfoResponse);
            }
            if (this.redeemGiftCardResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(26, this.redeemGiftCardResponse);
            }
            if (this.modifyLibraryResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(27, this.modifyLibraryResponse);
            }
            if (this.uploadDeviceConfigResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(28, this.uploadDeviceConfigResponse);
            }
            if (this.plusProfileResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(29, this.plusProfileResponse);
            }
            if (this.consumePurchaseResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(30, this.consumePurchaseResponse);
            }
            if (this.billingProfileResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(31, this.billingProfileResponse);
            }
            if (this.preparePurchaseResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(32, this.preparePurchaseResponse);
            }
            if (this.commitPurchaseResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(33, this.commitPurchaseResponse);
            }
            if (this.debugSettingsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(34, this.debugSettingsResponse);
            }
            if (this.checkIabPromoResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(35, this.checkIabPromoResponse);
            }
            if (this.userActivitySettingsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(36, this.userActivitySettingsResponse);
            }
            if (this.recordUserActivityResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(37, this.recordUserActivityResponse);
            }
            if (this.redeemCodeResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(38, this.redeemCodeResponse);
            }
            if (this.selfUpdateResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(39, this.selfUpdateResponse);
            }
            if (this.searchSuggestResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(40, this.searchSuggestResponse);
            }
            if (this.getInitialInstrumentFlowStateResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(41, this.getInitialInstrumentFlowStateResponse);
            }
            if (this.createInstrumentResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(42, this.createInstrumentResponse);
            }
            if (this.challengeResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(43, this.challengeResponse);
            }
            if (this.backupDeviceChoicesResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(44, this.backupDeviceChoicesResponse);
            }
            if (this.backupDocumentChoicesResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(45, this.backupDocumentChoicesResponse);
            }
            if (this.earlyUpdateResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(46, this.earlyUpdateResponse);
            }
            if (this.preloadsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(47, this.preloadsResponse);
            }
            if (this.myAccountResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(48, this.myAccountResponse);
            }
            if (this.contentFilterResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(49, this.contentFilterResponse);
            }
            if (this.experimentsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(50, this.experimentsResponse);
            }
            if (this.surveyResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(51, this.surveyResponse);
            }
            if (this.pingResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(52, this.pingResponse);
            }
            if (this.updateUserSettingResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(53, this.updateUserSettingResponse);
            }
            if (this.getUserSettingsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(54, this.getUserSettingsResponse);
            }
            if (this.getSharingSettingsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(56, this.getSharingSettingsResponse);
            }
            if (this.updateSharingSettingsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(57, this.updateSharingSettingsResponse);
            }
            if (this.reviewSnippetsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(58, this.reviewSnippetsResponse);
            }
            if (this.documentSharingStateResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(59, this.documentSharingStateResponse);
            }
            if (this.getFamilyPurchaseSettingResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(60, this.getFamilyPurchaseSettingResponse);
            }
            if (this.setFamilyPurchaseSettingReponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(61, this.setFamilyPurchaseSettingReponse);
            }
            if (this.dismissSurveyResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(62, this.dismissSurveyResponse);
            }
            if (this.instantPurchaseResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(63, this.instantPurchaseResponse);
            }
            if (this.familyFopResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(64, this.familyFopResponse);
            }
            if (this.monetaryGiftDetailsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(65, this.monetaryGiftDetailsResponse);
            }
            if (this.giftDialogDetailsResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(66, this.giftDialogDetailsResponse);
            }
            if (this.inAppPurchaseHistoryResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(67, this.inAppPurchaseHistoryResponse);
            }
            if (this.prepareUserRefundResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(68, this.prepareUserRefundResponse);
            }
            if (this.commitUserRefundResponse != null) {
                size += CodedOutputByteBufferNano.computeMessageSize(69, this.commitUserRefundResponse);
            }
            if (this.moduleDeliveryResponse != null) {
                return size + CodedOutputByteBufferNano.computeMessageSize(70, this.moduleDeliveryResponse);
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
                        if (this.listResponse == null) {
                            this.listResponse = new ListResponse();
                        }
                        x0.readMessage(this.listResponse);
                        break;
                    case 18:
                        if (this.detailsResponse == null) {
                            this.detailsResponse = new Details.DetailsResponse();
                        }
                        x0.readMessage(this.detailsResponse);
                        break;
                    case 26:
                        if (this.reviewResponse == null) {
                            this.reviewResponse = new ReviewResponse();
                        }
                        x0.readMessage(this.reviewResponse);
                        break;
                    case 34:
                        if (this.buyResponse == null) {
                            this.buyResponse = new Buy.BuyResponse();
                        }
                        x0.readMessage(this.buyResponse);
                        break;
                    case 42:
                        if (this.searchResponse == null) {
                            this.searchResponse = new Search.SearchResponse();
                        }
                        x0.readMessage(this.searchResponse);
                        break;
                    case 50:
                        if (this.tocResponse == null) {
                            this.tocResponse = new Toc.TocResponse();
                        }
                        x0.readMessage(this.tocResponse);
                        break;
                    case 58:
                        if (this.browseResponse == null) {
                            this.browseResponse = new Browse.BrowseResponse();
                        }
                        x0.readMessage(this.browseResponse);
                        break;
                    case 66:
                        if (this.purchaseStatusResponse == null) {
                            this.purchaseStatusResponse = new Buy.PurchaseStatusResponse();
                        }
                        x0.readMessage(this.purchaseStatusResponse);
                        break;
                    case 74:
                        if (this.updateInstrumentResponse == null) {
                            this.updateInstrumentResponse = new BuyInstruments.UpdateInstrumentResponse();
                        }
                        x0.readMessage(this.updateInstrumentResponse);
                        break;
                    case 82:
                        if (this.logResponse == null) {
                            this.logResponse = new Log.LogResponse();
                        }
                        x0.readMessage(this.logResponse);
                        break;
                    case 90:
                        if (this.checkInstrumentResponse == null) {
                            this.checkInstrumentResponse = new BuyInstruments.CheckInstrumentResponse();
                        }
                        x0.readMessage(this.checkInstrumentResponse);
                        break;
                    case 98:
                        if (this.plusOneResponse == null) {
                            this.plusOneResponse = new PlusOneResponse();
                        }
                        x0.readMessage(this.plusOneResponse);
                        break;
                    case 106:
                        if (this.flagContentResponse == null) {
                            this.flagContentResponse = new ContentFlagging.FlagContentResponse();
                        }
                        x0.readMessage(this.flagContentResponse);
                        break;
                    case 114:
                        if (this.ackNotificationResponse == null) {
                            this.ackNotificationResponse = new AckNotificationResponse();
                        }
                        x0.readMessage(this.ackNotificationResponse);
                        break;
                    case 122:
                        if (this.initiateAssociationResponse == null) {
                            this.initiateAssociationResponse = new CarrierBilling.InitiateAssociationResponse();
                        }
                        x0.readMessage(this.initiateAssociationResponse);
                        break;
                    case 130:
                        if (this.verifyAssociationResponse == null) {
                            this.verifyAssociationResponse = new CarrierBilling.VerifyAssociationResponse();
                        }
                        x0.readMessage(this.verifyAssociationResponse);
                        break;
                    case 138:
                        if (this.libraryReplicationResponse == null) {
                            this.libraryReplicationResponse = new LibraryReplicationResponse();
                        }
                        x0.readMessage(this.libraryReplicationResponse);
                        break;
                    case 146:
                        if (this.revokeResponse == null) {
                            this.revokeResponse = new RevokeResponse();
                        }
                        x0.readMessage(this.revokeResponse);
                        break;
                    case 154:
                        if (this.bulkDetailsResponse == null) {
                            this.bulkDetailsResponse = new Details.BulkDetailsResponse();
                        }
                        x0.readMessage(this.bulkDetailsResponse);
                        break;
                    case 162:
                        if (this.resolveLinkResponse == null) {
                            this.resolveLinkResponse = new ResolveLink.ResolvedLink();
                        }
                        x0.readMessage(this.resolveLinkResponse);
                        break;
                    case 170:
                        if (this.deliveryResponse == null) {
                            this.deliveryResponse = new DeliveryResponse();
                        }
                        x0.readMessage(this.deliveryResponse);
                        break;
                    case 178:
                        if (this.acceptTosResponse == null) {
                            this.acceptTosResponse = new AcceptTosResponse();
                        }
                        x0.readMessage(this.acceptTosResponse);
                        break;
                    case 186:
                        if (this.rateSuggestedContentResponse == null) {
                            this.rateSuggestedContentResponse = new RateSuggestedContentResponse();
                        }
                        x0.readMessage(this.rateSuggestedContentResponse);
                        break;
                    case 194:
                        if (this.checkPromoOfferResponse == null) {
                            this.checkPromoOfferResponse = new CheckPromoOfferResponse();
                        }
                        x0.readMessage(this.checkPromoOfferResponse);
                        break;
                    case 202:
                        if (this.instrumentSetupInfoResponse == null) {
                            this.instrumentSetupInfoResponse = new BuyInstruments.InstrumentSetupInfoResponse();
                        }
                        x0.readMessage(this.instrumentSetupInfoResponse);
                        break;
                    case 210:
                        if (this.redeemGiftCardResponse == null) {
                            this.redeemGiftCardResponse = new BuyInstruments.RedeemGiftCardResponse();
                        }
                        x0.readMessage(this.redeemGiftCardResponse);
                        break;
                    case 218:
                        if (this.modifyLibraryResponse == null) {
                            this.modifyLibraryResponse = new ModifyLibrary.ModifyLibraryResponse();
                        }
                        x0.readMessage(this.modifyLibraryResponse);
                        break;
                    case 226:
                        if (this.uploadDeviceConfigResponse == null) {
                            this.uploadDeviceConfigResponse = new UploadDeviceConfig.UploadDeviceConfigResponse();
                        }
                        x0.readMessage(this.uploadDeviceConfigResponse);
                        break;
                    case 234:
                        if (this.plusProfileResponse == null) {
                            this.plusProfileResponse = new PlusProfileResponse();
                        }
                        x0.readMessage(this.plusProfileResponse);
                        break;
                    case 242:
                        if (this.consumePurchaseResponse == null) {
                            this.consumePurchaseResponse = new ConsumePurchaseResponse();
                        }
                        x0.readMessage(this.consumePurchaseResponse);
                        break;
                    case 250:
                        if (this.billingProfileResponse == null) {
                            this.billingProfileResponse = new BuyInstruments.BillingProfileResponse();
                        }
                        x0.readMessage(this.billingProfileResponse);
                        break;
                    case 258:
                        if (this.preparePurchaseResponse == null) {
                            this.preparePurchaseResponse = new Purchase.PreparePurchaseResponse();
                        }
                        x0.readMessage(this.preparePurchaseResponse);
                        break;
                    case 266:
                        if (this.commitPurchaseResponse == null) {
                            this.commitPurchaseResponse = new Purchase.CommitPurchaseResponse();
                        }
                        x0.readMessage(this.commitPurchaseResponse);
                        break;
                    case 274:
                        if (this.debugSettingsResponse == null) {
                            this.debugSettingsResponse = new DebugSettingsResponse();
                        }
                        x0.readMessage(this.debugSettingsResponse);
                        break;
                    case 282:
                        if (this.checkIabPromoResponse == null) {
                            this.checkIabPromoResponse = new BuyInstruments.CheckIabPromoResponse();
                        }
                        x0.readMessage(this.checkIabPromoResponse);
                        break;
                    case 290:
                        if (this.userActivitySettingsResponse == null) {
                            this.userActivitySettingsResponse = new UserActivity.UserActivitySettingsResponse();
                        }
                        x0.readMessage(this.userActivitySettingsResponse);
                        break;
                    case 298:
                        if (this.recordUserActivityResponse == null) {
                            this.recordUserActivityResponse = new UserActivity.RecordUserActivityResponse();
                        }
                        x0.readMessage(this.recordUserActivityResponse);
                        break;
                    case 306:
                        if (this.redeemCodeResponse == null) {
                            this.redeemCodeResponse = new PromoCode.RedeemCodeResponse();
                        }
                        x0.readMessage(this.redeemCodeResponse);
                        break;
                    case 314:
                        if (this.selfUpdateResponse == null) {
                            this.selfUpdateResponse = new SelfUpdateResponse();
                        }
                        x0.readMessage(this.selfUpdateResponse);
                        break;
                    case 322:
                        if (this.searchSuggestResponse == null) {
                            this.searchSuggestResponse = new SearchSuggest.SearchSuggestResponse();
                        }
                        x0.readMessage(this.searchSuggestResponse);
                        break;
                    case 330:
                        if (this.getInitialInstrumentFlowStateResponse == null) {
                            this.getInitialInstrumentFlowStateResponse = new BuyInstruments.GetInitialInstrumentFlowStateResponse();
                        }
                        x0.readMessage(this.getInitialInstrumentFlowStateResponse);
                        break;
                    case 338:
                        if (this.createInstrumentResponse == null) {
                            this.createInstrumentResponse = new BuyInstruments.CreateInstrumentResponse();
                        }
                        x0.readMessage(this.createInstrumentResponse);
                        break;
                    case 346:
                        if (this.challengeResponse == null) {
                            this.challengeResponse = new ChallengeResponse();
                        }
                        x0.readMessage(this.challengeResponse);
                        break;
                    case 354:
                        if (this.backupDeviceChoicesResponse == null) {
                            this.backupDeviceChoicesResponse = new Restore.GetBackupDeviceChoicesResponse();
                        }
                        x0.readMessage(this.backupDeviceChoicesResponse);
                        break;
                    case 362:
                        if (this.backupDocumentChoicesResponse == null) {
                            this.backupDocumentChoicesResponse = new Restore.GetBackupDocumentChoicesResponse();
                        }
                        x0.readMessage(this.backupDocumentChoicesResponse);
                        break;
                    case 370:
                        if (this.earlyUpdateResponse == null) {
                            this.earlyUpdateResponse = new EarlyUpdateResponse();
                        }
                        x0.readMessage(this.earlyUpdateResponse);
                        break;
                    case 378:
                        if (this.preloadsResponse == null) {
                            this.preloadsResponse = new Preloads.PreloadsResponse();
                        }
                        x0.readMessage(this.preloadsResponse);
                        break;
                    case 386:
                        if (this.myAccountResponse == null) {
                            this.myAccountResponse = new MyAccountResponse();
                        }
                        x0.readMessage(this.myAccountResponse);
                        break;
                    case 394:
                        if (this.contentFilterResponse == null) {
                            this.contentFilterResponse = new ContentFilters.ContentFilterSettingsResponse();
                        }
                        x0.readMessage(this.contentFilterResponse);
                        break;
                    case 402:
                        if (this.experimentsResponse == null) {
                            this.experimentsResponse = new ExperimentsResponse();
                        }
                        x0.readMessage(this.experimentsResponse);
                        break;
                    case 410:
                        if (this.surveyResponse == null) {
                            this.surveyResponse = new SurveyResponse();
                        }
                        x0.readMessage(this.surveyResponse);
                        break;
                    case 418:
                        if (this.pingResponse == null) {
                            this.pingResponse = new PingResponse();
                        }
                        x0.readMessage(this.pingResponse);
                        break;
                    case 426:
                        if (this.updateUserSettingResponse == null) {
                            this.updateUserSettingResponse = new UpdateUserSettingResponse();
                        }
                        x0.readMessage(this.updateUserSettingResponse);
                        break;
                    case 434:
                        if (this.getUserSettingsResponse == null) {
                            this.getUserSettingsResponse = new GetUserSettingsResponse();
                        }
                        x0.readMessage(this.getUserSettingsResponse);
                        break;
                    case 450:
                        if (this.getSharingSettingsResponse == null) {
                            this.getSharingSettingsResponse = new SharingSetting.GetFamilySharingSettingsResponse();
                        }
                        x0.readMessage(this.getSharingSettingsResponse);
                        break;
                    case 458:
                        if (this.updateSharingSettingsResponse == null) {
                            this.updateSharingSettingsResponse = new SharingSetting.UpdateFamilySharingSettingsResponse();
                        }
                        x0.readMessage(this.updateSharingSettingsResponse);
                        break;
                    case 466:
                        if (this.reviewSnippetsResponse == null) {
                            this.reviewSnippetsResponse = new ReviewSnippetsResponse();
                        }
                        x0.readMessage(this.reviewSnippetsResponse);
                        break;
                    case 474:
                        if (this.documentSharingStateResponse == null) {
                            this.documentSharingStateResponse = new DocumentSharingStateResponse();
                        }
                        x0.readMessage(this.documentSharingStateResponse);
                        break;
                    case 482:
                        if (this.getFamilyPurchaseSettingResponse == null) {
                            this.getFamilyPurchaseSettingResponse = new GetFamilyPurchaseSettingResponse();
                        }
                        x0.readMessage(this.getFamilyPurchaseSettingResponse);
                        break;
                    case 490:
                        if (this.setFamilyPurchaseSettingReponse == null) {
                            this.setFamilyPurchaseSettingReponse = new SetFamilyPurchaseSettingResponse();
                        }
                        x0.readMessage(this.setFamilyPurchaseSettingReponse);
                        break;
                    case 498:
                        if (this.dismissSurveyResponse == null) {
                            this.dismissSurveyResponse = new DismissSurveyResponse();
                        }
                        x0.readMessage(this.dismissSurveyResponse);
                        break;
                    case 506:
                        if (this.instantPurchaseResponse == null) {
                            this.instantPurchaseResponse = new Purchase.InstantPurchaseResponse();
                        }
                        x0.readMessage(this.instantPurchaseResponse);
                        break;
                    case 514:
                        if (this.familyFopResponse == null) {
                            this.familyFopResponse = new FamilyFopResponse();
                        }
                        x0.readMessage(this.familyFopResponse);
                        break;
                    case 522:
                        if (this.monetaryGiftDetailsResponse == null) {
                            this.monetaryGiftDetailsResponse = new MonetaryGiftDetailsResponse();
                        }
                        x0.readMessage(this.monetaryGiftDetailsResponse);
                        break;
                    case 530:
                        if (this.giftDialogDetailsResponse == null) {
                            this.giftDialogDetailsResponse = new GiftDialogDetailsResponse();
                        }
                        x0.readMessage(this.giftDialogDetailsResponse);
                        break;
                    case 538:
                        if (this.inAppPurchaseHistoryResponse == null) {
                            this.inAppPurchaseHistoryResponse = new Purchase.InAppPurchaseHistoryResponse();
                        }
                        x0.readMessage(this.inAppPurchaseHistoryResponse);
                        break;
                    case 546:
                        if (this.prepareUserRefundResponse == null) {
                            this.prepareUserRefundResponse = new UserRefund.PrepareUserRefundResponse();
                        }
                        x0.readMessage(this.prepareUserRefundResponse);
                        break;
                    case 554:
                        if (this.commitUserRefundResponse == null) {
                            this.commitUserRefundResponse = new UserRefund.CommitUserRefundResponse();
                        }
                        x0.readMessage(this.commitUserRefundResponse);
                        break;
                    case 562:
                        if (this.moduleDeliveryResponse == null) {
                            this.moduleDeliveryResponse = new ModuleDeliveryResponse();
                        }
                        x0.readMessage(this.moduleDeliveryResponse);
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
}
