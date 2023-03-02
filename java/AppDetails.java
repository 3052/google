package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class AppDetails extends MessageNano {
    public String developerName = "";
    public boolean hasDeveloperName = false;
    public int majorVersionNumber = 0;
    public boolean hasMajorVersionNumber = false;
    public int versionCode = 0;
    public boolean hasVersionCode = false;
    public String versionString = "";
    public boolean hasVersionString = false;
    public String title = "";
    public boolean hasTitle = false;
    public String[] appCategory = WireFormatNano.EMPTY_STRING_ARRAY;
    public int contentRating = 0;
    public boolean hasContentRating = false;
    public long installationSize = 0;
    public boolean hasInstallationSize = false;
    public String[] permission = WireFormatNano.EMPTY_STRING_ARRAY;
    public String developerEmail = "";
    public boolean hasDeveloperEmail = false;
    public String developerWebsite = "";
    public boolean hasDeveloperWebsite = false;
    public String numDownloads = "";
    public boolean hasNumDownloads = false;
    public String packageName = "";
    public boolean hasPackageName = false;
    public String recentChangesHtml = "";
    public boolean hasRecentChangesHtml = false;
    public String uploadDate = "";
    public boolean hasUploadDate = false;
    public FileMetadata[] file = FileMetadata.emptyArray();
    public String appType = "";
    public boolean hasAppType = false;
    public CertificateSet[] certificateSet = CertificateSet.emptyArray();
    public String[] certificateHash = WireFormatNano.EMPTY_STRING_ARRAY;
    public boolean variesByAccount = true;
    public boolean hasVariesByAccount = false;
    public String[] autoAcquireFreeAppIfHigherVersionAvailableTag = WireFormatNano.EMPTY_STRING_ARRAY;
    public boolean declaresIab = false;
    public boolean hasDeclaresIab = false;
    public String[] splitId = WireFormatNano.EMPTY_STRING_ARRAY;
    public boolean gamepadRequired = false;
    public boolean hasGamepadRequired = false;
    public boolean externallyHosted = false;
    public boolean hasExternallyHosted = false;
    public boolean everExternallyHosted = false;
    public boolean hasEverExternallyHosted = false;
    public String installNotes = "";
    public boolean hasInstallNotes = false;
    public int installLocation = 0;
    public boolean hasInstallLocation = false;
    public int targetSdkVersion = 0;
    public boolean hasTargetSdkVersion = false;
    public String preregistrationPromoCode = "";
    public boolean hasPreregistrationPromoCode = false;
    public InstallDetails installDetails = null;

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
                    this.developerName = x0.readString();
                    this.hasDeveloperName = true;
                    break;
                case 16:
                    this.majorVersionNumber = x0.readRawVarint32();
                    this.hasMajorVersionNumber = true;
                    break;
                case 24:
                    this.versionCode = x0.readRawVarint32();
                    this.hasVersionCode = true;
                    break;
                case 34:
                    this.versionString = x0.readString();
                    this.hasVersionString = true;
                    break;
                case 42:
                    this.title = x0.readString();
                    this.hasTitle = true;
                    break;
                case 58:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 58);
                    int length3 = this.appCategory == null ? 0 : this.appCategory.length;
                    String[] strArr = new String[repeatedFieldArrayLength + length3];
                    if (length3 != 0) {
                        System.arraycopy(this.appCategory, 0, strArr, 0, length3);
                    }
                    while (length3 < strArr.length - 1) {
                        strArr[length3] = x0.readString();
                        x0.readTag();
                        length3++;
                    }
                    strArr[length3] = x0.readString();
                    this.appCategory = strArr;
                    break;
                case 64:
                    this.contentRating = x0.readRawVarint32();
                    this.hasContentRating = true;
                    break;
                case 72:
                    this.installationSize = x0.readRawVarint64();
                    this.hasInstallationSize = true;
                    break;
                case 82:
                    int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 82);
                    int length4 = this.permission == null ? 0 : this.permission.length;
                    String[] strArr2 = new String[repeatedFieldArrayLength2 + length4];
                    if (length4 != 0) {
                        System.arraycopy(this.permission, 0, strArr2, 0, length4);
                    }
                    while (length4 < strArr2.length - 1) {
                        strArr2[length4] = x0.readString();
                        x0.readTag();
                        length4++;
                    }
                    strArr2[length4] = x0.readString();
                    this.permission = strArr2;
                    break;
                case 90:
                    this.developerEmail = x0.readString();
                    this.hasDeveloperEmail = true;
                    break;
                case 98:
                    this.developerWebsite = x0.readString();
                    this.hasDeveloperWebsite = true;
                    break;
                case 106:
                    this.numDownloads = x0.readString();
                    this.hasNumDownloads = true;
                    break;
                case 114:
                    this.packageName = x0.readString();
                    this.hasPackageName = true;
                    break;
                case 122:
                    this.recentChangesHtml = x0.readString();
                    this.hasRecentChangesHtml = true;
                    break;
                case 130:
                    this.uploadDate = x0.readString();
                    this.hasUploadDate = true;
                    break;
                case 138:
                    int repeatedFieldArrayLength3 = WireFormatNano.getRepeatedFieldArrayLength(x0, 138);
                    if (this.file == null) {
                        length2 = 0;
                    } else {
                        length2 = this.file.length;
                    }
                    FileMetadata[] fileMetadataArr = new FileMetadata[repeatedFieldArrayLength3 + length2];
                    if (length2 != 0) {
                        System.arraycopy(this.file, 0, fileMetadataArr, 0, length2);
                    }
                    while (length2 < fileMetadataArr.length - 1) {
                        fileMetadataArr[length2] = new FileMetadata();
                        x0.readMessage(fileMetadataArr[length2]);
                        x0.readTag();
                        length2++;
                    }
                    fileMetadataArr[length2] = new FileMetadata();
                    x0.readMessage(fileMetadataArr[length2]);
                    this.file = fileMetadataArr;
                    break;
                case 146:
                    this.appType = x0.readString();
                    this.hasAppType = true;
                    break;
                case 154:
                    int repeatedFieldArrayLength4 = WireFormatNano.getRepeatedFieldArrayLength(x0, 154);
                    int length5 = this.certificateHash == null ? 0 : this.certificateHash.length;
                    String[] strArr3 = new String[repeatedFieldArrayLength4 + length5];
                    if (length5 != 0) {
                        System.arraycopy(this.certificateHash, 0, strArr3, 0, length5);
                    }
                    while (length5 < strArr3.length - 1) {
                        strArr3[length5] = x0.readString();
                        x0.readTag();
                        length5++;
                    }
                    strArr3[length5] = x0.readString();
                    this.certificateHash = strArr3;
                    break;
                case 168:
                    this.variesByAccount = x0.readBool();
                    this.hasVariesByAccount = true;
                    break;
                case 178:
                    int repeatedFieldArrayLength5 = WireFormatNano.getRepeatedFieldArrayLength(x0, 178);
                    if (this.certificateSet == null) {
                        length = 0;
                    } else {
                        length = this.certificateSet.length;
                    }
                    CertificateSet[] certificateSetArr = new CertificateSet[repeatedFieldArrayLength5 + length];
                    if (length != 0) {
                        System.arraycopy(this.certificateSet, 0, certificateSetArr, 0, length);
                    }
                    while (length < certificateSetArr.length - 1) {
                        certificateSetArr[length] = new CertificateSet();
                        x0.readMessage(certificateSetArr[length]);
                        x0.readTag();
                        length++;
                    }
                    certificateSetArr[length] = new CertificateSet();
                    x0.readMessage(certificateSetArr[length]);
                    this.certificateSet = certificateSetArr;
                    break;
                case 186:
                    int repeatedFieldArrayLength6 = WireFormatNano.getRepeatedFieldArrayLength(x0, 186);
                    int length6 = this.autoAcquireFreeAppIfHigherVersionAvailableTag == null ? 0 : this.autoAcquireFreeAppIfHigherVersionAvailableTag.length;
                    String[] strArr4 = new String[repeatedFieldArrayLength6 + length6];
                    if (length6 != 0) {
                        System.arraycopy(this.autoAcquireFreeAppIfHigherVersionAvailableTag, 0, strArr4, 0, length6);
                    }
                    while (length6 < strArr4.length - 1) {
                        strArr4[length6] = x0.readString();
                        x0.readTag();
                        length6++;
                    }
                    strArr4[length6] = x0.readString();
                    this.autoAcquireFreeAppIfHigherVersionAvailableTag = strArr4;
                    break;
                case 192:
                    this.declaresIab = x0.readBool();
                    this.hasDeclaresIab = true;
                    break;
                case 202:
                    int repeatedFieldArrayLength7 = WireFormatNano.getRepeatedFieldArrayLength(x0, 202);
                    int length7 = this.splitId == null ? 0 : this.splitId.length;
                    String[] strArr5 = new String[repeatedFieldArrayLength7 + length7];
                    if (length7 != 0) {
                        System.arraycopy(this.splitId, 0, strArr5, 0, length7);
                    }
                    while (length7 < strArr5.length - 1) {
                        strArr5[length7] = x0.readString();
                        x0.readTag();
                        length7++;
                    }
                    strArr5[length7] = x0.readString();
                    this.splitId = strArr5;
                    break;
                case 208:
                    this.gamepadRequired = x0.readBool();
                    this.hasGamepadRequired = true;
                    break;
                case 216:
                    this.externallyHosted = x0.readBool();
                    this.hasExternallyHosted = true;
                    break;
                case 224:
                    this.everExternallyHosted = x0.readBool();
                    this.hasEverExternallyHosted = true;
                    break;
                case 242:
                    this.installNotes = x0.readString();
                    this.hasInstallNotes = true;
                    break;
                case 248:
                    int readRawVarint32 = x0.readRawVarint32();
                    switch (readRawVarint32) {
                        case 0:
                        case 1:
                        case 2:
                        case 3:
                            this.installLocation = readRawVarint32;
                            this.hasInstallLocation = true;
                            continue;
                    }
                case 256:
                    this.targetSdkVersion = x0.readRawVarint32();
                    this.hasTargetSdkVersion = true;
                    break;
                case 266:
                    this.preregistrationPromoCode = x0.readString();
                    this.hasPreregistrationPromoCode = true;
                    break;
                case 274:
                    if (this.installDetails == null) {
                        this.installDetails = new InstallDetails();
                    }
                    x0.readMessage(this.installDetails);
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

    public AppDetails() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasDeveloperName || !this.developerName.equals("")) {
            output.writeString(1, this.developerName);
        }
        if (this.hasMajorVersionNumber || this.majorVersionNumber != 0) {
            output.writeInt32(2, this.majorVersionNumber);
        }
        if (this.hasVersionCode || this.versionCode != 0) {
            output.writeInt32(3, this.versionCode);
        }
        if (this.hasVersionString || !this.versionString.equals("")) {
            output.writeString(4, this.versionString);
        }
        if (this.hasTitle || !this.title.equals("")) {
            output.writeString(5, this.title);
        }
        if (this.appCategory != null && this.appCategory.length > 0) {
            for (int i = 0; i < this.appCategory.length; i++) {
                String element = this.appCategory[i];
                if (element != null) {
                    output.writeString(7, element);
                }
            }
        }
        if (this.hasContentRating || this.contentRating != 0) {
            output.writeInt32(8, this.contentRating);
        }
        if (this.hasInstallationSize || this.installationSize != 0) {
            output.writeInt64(9, this.installationSize);
        }
        if (this.permission != null && this.permission.length > 0) {
            for (int i2 = 0; i2 < this.permission.length; i2++) {
                String element2 = this.permission[i2];
                if (element2 != null) {
                    output.writeString(10, element2);
                }
            }
        }
        if (this.hasDeveloperEmail || !this.developerEmail.equals("")) {
            output.writeString(11, this.developerEmail);
        }
        if (this.hasDeveloperWebsite || !this.developerWebsite.equals("")) {
            output.writeString(12, this.developerWebsite);
        }
        if (this.hasNumDownloads || !this.numDownloads.equals("")) {
            output.writeString(13, this.numDownloads);
        }
        if (this.hasPackageName || !this.packageName.equals("")) {
            output.writeString(14, this.packageName);
        }
        if (this.hasRecentChangesHtml || !this.recentChangesHtml.equals("")) {
            output.writeString(15, this.recentChangesHtml);
        }
        if (this.hasUploadDate || !this.uploadDate.equals("")) {
            output.writeString(16, this.uploadDate);
        }
        if (this.file != null && this.file.length > 0) {
            for (int i3 = 0; i3 < this.file.length; i3++) {
                FileMetadata element3 = this.file[i3];
                if (element3 != null) {
                    output.writeMessage(17, element3);
                }
            }
        }
        if (this.hasAppType || !this.appType.equals("")) {
            output.writeString(18, this.appType);
        }
        if (this.certificateHash != null && this.certificateHash.length > 0) {
            for (int i4 = 0; i4 < this.certificateHash.length; i4++) {
                String element4 = this.certificateHash[i4];
                if (element4 != null) {
                    output.writeString(19, element4);
                }
            }
        }
        if (this.hasVariesByAccount || !this.variesByAccount) {
            output.writeBool(21, this.variesByAccount);
        }
        if (this.certificateSet != null && this.certificateSet.length > 0) {
            for (int i5 = 0; i5 < this.certificateSet.length; i5++) {
                CertificateSet element5 = this.certificateSet[i5];
                if (element5 != null) {
                    output.writeMessage(22, element5);
                }
            }
        }
        if (this.autoAcquireFreeAppIfHigherVersionAvailableTag != null && this.autoAcquireFreeAppIfHigherVersionAvailableTag.length > 0) {
            for (int i6 = 0; i6 < this.autoAcquireFreeAppIfHigherVersionAvailableTag.length; i6++) {
                String element6 = this.autoAcquireFreeAppIfHigherVersionAvailableTag[i6];
                if (element6 != null) {
                    output.writeString(23, element6);
                }
            }
        }
        if (this.hasDeclaresIab || this.declaresIab) {
            output.writeBool(24, this.declaresIab);
        }
        if (this.splitId != null && this.splitId.length > 0) {
            for (int i7 = 0; i7 < this.splitId.length; i7++) {
                String element7 = this.splitId[i7];
                if (element7 != null) {
                    output.writeString(25, element7);
                }
            }
        }
        if (this.hasGamepadRequired || this.gamepadRequired) {
            output.writeBool(26, this.gamepadRequired);
        }
        if (this.hasExternallyHosted || this.externallyHosted) {
            output.writeBool(27, this.externallyHosted);
        }
        if (this.hasEverExternallyHosted || this.everExternallyHosted) {
            output.writeBool(28, this.everExternallyHosted);
        }
        if (this.hasInstallNotes || !this.installNotes.equals("")) {
            output.writeString(30, this.installNotes);
        }
        if (this.installLocation != 0 || this.hasInstallLocation) {
            output.writeInt32(31, this.installLocation);
        }
        if (this.hasTargetSdkVersion || this.targetSdkVersion != 0) {
            output.writeInt32(32, this.targetSdkVersion);
        }
        if (this.hasPreregistrationPromoCode || !this.preregistrationPromoCode.equals("")) {
            output.writeString(33, this.preregistrationPromoCode);
        }
        if (this.installDetails != null) {
            output.writeMessage(34, this.installDetails);
        }
        super.writeTo(output);
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasDeveloperName || !this.developerName.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(1, this.developerName);
        }
        if (this.hasMajorVersionNumber || this.majorVersionNumber != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(2, this.majorVersionNumber);
        }
        if (this.hasVersionCode || this.versionCode != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(3, this.versionCode);
        }
        if (this.hasVersionString || !this.versionString.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(4, this.versionString);
        }
        if (this.hasTitle || !this.title.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(5, this.title);
        }
        if (this.appCategory != null && this.appCategory.length > 0) {
            int dataCount = 0;
            int dataSize = 0;
            for (int i = 0; i < this.appCategory.length; i++) {
                String element = this.appCategory[i];
                if (element != null) {
                    dataCount++;
                    dataSize += CodedOutputByteBufferNano.computeStringSizeNoTag(element);
                }
            }
            size = size + dataSize + (dataCount * 1);
        }
        if (this.hasContentRating || this.contentRating != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(8, this.contentRating);
        }
        if (this.hasInstallationSize || this.installationSize != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(9, this.installationSize);
        }
        if (this.permission != null && this.permission.length > 0) {
            int dataCount2 = 0;
            int dataSize2 = 0;
            for (int i2 = 0; i2 < this.permission.length; i2++) {
                String element2 = this.permission[i2];
                if (element2 != null) {
                    dataCount2++;
                    dataSize2 += CodedOutputByteBufferNano.computeStringSizeNoTag(element2);
                }
            }
            size = size + dataSize2 + (dataCount2 * 1);
        }
        if (this.hasDeveloperEmail || !this.developerEmail.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(11, this.developerEmail);
        }
        if (this.hasDeveloperWebsite || !this.developerWebsite.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(12, this.developerWebsite);
        }
        if (this.hasNumDownloads || !this.numDownloads.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(13, this.numDownloads);
        }
        if (this.hasPackageName || !this.packageName.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(14, this.packageName);
        }
        if (this.hasRecentChangesHtml || !this.recentChangesHtml.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(15, this.recentChangesHtml);
        }
        if (this.hasUploadDate || !this.uploadDate.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(16, this.uploadDate);
        }
        if (this.file != null && this.file.length > 0) {
            for (int i3 = 0; i3 < this.file.length; i3++) {
                FileMetadata element3 = this.file[i3];
                if (element3 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(17, element3);
                }
            }
        }
        if (this.hasAppType || !this.appType.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(18, this.appType);
        }
        if (this.certificateHash != null && this.certificateHash.length > 0) {
            int dataCount3 = 0;
            int dataSize3 = 0;
            for (int i4 = 0; i4 < this.certificateHash.length; i4++) {
                String element4 = this.certificateHash[i4];
                if (element4 != null) {
                    dataCount3++;
                    dataSize3 += CodedOutputByteBufferNano.computeStringSizeNoTag(element4);
                }
            }
            size = size + dataSize3 + (dataCount3 * 2);
        }
        if (this.hasVariesByAccount || !this.variesByAccount) {
            size += CodedOutputByteBufferNano.computeTagSize(21) + 1;
        }
        if (this.certificateSet != null && this.certificateSet.length > 0) {
            for (int i5 = 0; i5 < this.certificateSet.length; i5++) {
                CertificateSet element5 = this.certificateSet[i5];
                if (element5 != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(22, element5);
                }
            }
        }
        if (this.autoAcquireFreeAppIfHigherVersionAvailableTag != null && this.autoAcquireFreeAppIfHigherVersionAvailableTag.length > 0) {
            int dataCount4 = 0;
            int dataSize4 = 0;
            for (int i6 = 0; i6 < this.autoAcquireFreeAppIfHigherVersionAvailableTag.length; i6++) {
                String element6 = this.autoAcquireFreeAppIfHigherVersionAvailableTag[i6];
                if (element6 != null) {
                    dataCount4++;
                    dataSize4 += CodedOutputByteBufferNano.computeStringSizeNoTag(element6);
                }
            }
            size = size + dataSize4 + (dataCount4 * 2);
        }
        if (this.hasDeclaresIab || this.declaresIab) {
            size += CodedOutputByteBufferNano.computeTagSize(24) + 1;
        }
        if (this.splitId != null && this.splitId.length > 0) {
            int dataCount5 = 0;
            int dataSize5 = 0;
            for (int i7 = 0; i7 < this.splitId.length; i7++) {
                String element7 = this.splitId[i7];
                if (element7 != null) {
                    dataCount5++;
                    dataSize5 += CodedOutputByteBufferNano.computeStringSizeNoTag(element7);
                }
            }
            size = size + dataSize5 + (dataCount5 * 2);
        }
        if (this.hasGamepadRequired || this.gamepadRequired) {
            size += CodedOutputByteBufferNano.computeTagSize(26) + 1;
        }
        if (this.hasExternallyHosted || this.externallyHosted) {
            size += CodedOutputByteBufferNano.computeTagSize(27) + 1;
        }
        if (this.hasEverExternallyHosted || this.everExternallyHosted) {
            size += CodedOutputByteBufferNano.computeTagSize(28) + 1;
        }
        if (this.hasInstallNotes || !this.installNotes.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(30, this.installNotes);
        }
        if (this.installLocation != 0 || this.hasInstallLocation) {
            size += CodedOutputByteBufferNano.computeInt32Size(31, this.installLocation);
        }
        if (this.hasTargetSdkVersion || this.targetSdkVersion != 0) {
            size += CodedOutputByteBufferNano.computeInt32Size(32, this.targetSdkVersion);
        }
        if (this.hasPreregistrationPromoCode || !this.preregistrationPromoCode.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(33, this.preregistrationPromoCode);
        }
        if (this.installDetails != null) {
            return size + CodedOutputByteBufferNano.computeMessageSize(34, this.installDetails);
        }
        return size;
    }
}
