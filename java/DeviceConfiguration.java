package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public interface DeviceConfiguration {

    /* loaded from: classes.dex */
    public static final class DeviceConfigurationProto extends MessageNano {
        public int touchScreen = 0;
        public boolean hasTouchScreen = false;
        public int keyboard = 0;
        public boolean hasKeyboard = false;
        public int navigation = 0;
        public boolean hasNavigation = false;
        public int screenLayout = 0;
        public boolean hasScreenLayout = false;
        public boolean hasHardKeyboard = false;
        public boolean hasHasHardKeyboard = false;
        public boolean hasFiveWayNavigation = false;
        public boolean hasHasFiveWayNavigation = false;
        public int screenDensity = 0;
        public boolean hasScreenDensity = false;
        public int screenWidth = 0;
        public boolean hasScreenWidth = false;
        public int screenHeight = 0;
        public boolean hasScreenHeight = false;
        public int glEsVersion = 0;
        public boolean hasGlEsVersion = false;
        public String[] systemSharedLibrary = WireFormatNano.EMPTY_STRING_ARRAY;
        public String[] systemAvailableFeature = WireFormatNano.EMPTY_STRING_ARRAY;
        public String[] nativePlatform = WireFormatNano.EMPTY_STRING_ARRAY;
        public String[] systemSupportedLocale = WireFormatNano.EMPTY_STRING_ARRAY;
        public String[] glExtension = WireFormatNano.EMPTY_STRING_ARRAY;
        public int maxApkDownloadSizeMb = 50;
        public boolean hasMaxApkDownloadSizeMb = false;
        public int smallestScreenWidthDp = 0;
        public boolean hasSmallestScreenWidthDp = false;
        public boolean lowRamDevice = false;
        public boolean hasLowRamDevice = false;
        public long totalMemoryBytes = 0;
        public boolean hasTotalMemoryBytes = false;
        public int maxNumOfCpuCores = 0;
        public boolean hasMaxNumOfCpuCores = false;

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
                                this.touchScreen = readRawVarint32;
                                this.hasTouchScreen = true;
                                continue;
                        }
                    case 16:
                        int readRawVarint322 = x0.readRawVarint32();
                        switch (readRawVarint322) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                                this.keyboard = readRawVarint322;
                                this.hasKeyboard = true;
                                continue;
                        }
                    case 24:
                        int readRawVarint323 = x0.readRawVarint32();
                        switch (readRawVarint323) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                                this.navigation = readRawVarint323;
                                this.hasNavigation = true;
                                continue;
                        }
                    case 32:
                        int readRawVarint324 = x0.readRawVarint32();
                        switch (readRawVarint324) {
                            case 0:
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                                this.screenLayout = readRawVarint324;
                                this.hasScreenLayout = true;
                                continue;
                        }
                    case 40:
                        this.hasHardKeyboard = x0.readBool();
                        this.hasHasHardKeyboard = true;
                        break;
                    case 48:
                        this.hasFiveWayNavigation = x0.readBool();
                        this.hasHasFiveWayNavigation = true;
                        break;
                    case 56:
                        this.screenDensity = x0.readRawVarint32();
                        this.hasScreenDensity = true;
                        break;
                    case 64:
                        this.glEsVersion = x0.readRawVarint32();
                        this.hasGlEsVersion = true;
                        break;
                    case 74:
                        int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 74);
                        int length = this.systemSharedLibrary == null ? 0 : this.systemSharedLibrary.length;
                        String[] strArr = new String[repeatedFieldArrayLength + length];
                        if (length != 0) {
                            System.arraycopy(this.systemSharedLibrary, 0, strArr, 0, length);
                        }
                        while (length < strArr.length - 1) {
                            strArr[length] = x0.readString();
                            x0.readTag();
                            length++;
                        }
                        strArr[length] = x0.readString();
                        this.systemSharedLibrary = strArr;
                        break;
                    case 82:
                        int repeatedFieldArrayLength2 = WireFormatNano.getRepeatedFieldArrayLength(x0, 82);
                        int length2 = this.systemAvailableFeature == null ? 0 : this.systemAvailableFeature.length;
                        String[] strArr2 = new String[repeatedFieldArrayLength2 + length2];
                        if (length2 != 0) {
                            System.arraycopy(this.systemAvailableFeature, 0, strArr2, 0, length2);
                        }
                        while (length2 < strArr2.length - 1) {
                            strArr2[length2] = x0.readString();
                            x0.readTag();
                            length2++;
                        }
                        strArr2[length2] = x0.readString();
                        this.systemAvailableFeature = strArr2;
                        break;
                    case 90:
                        int repeatedFieldArrayLength3 = WireFormatNano.getRepeatedFieldArrayLength(x0, 90);
                        int length3 = this.nativePlatform == null ? 0 : this.nativePlatform.length;
                        String[] strArr3 = new String[repeatedFieldArrayLength3 + length3];
                        if (length3 != 0) {
                            System.arraycopy(this.nativePlatform, 0, strArr3, 0, length3);
                        }
                        while (length3 < strArr3.length - 1) {
                            strArr3[length3] = x0.readString();
                            x0.readTag();
                            length3++;
                        }
                        strArr3[length3] = x0.readString();
                        this.nativePlatform = strArr3;
                        break;
                    case 96:
                        this.screenWidth = x0.readRawVarint32();
                        this.hasScreenWidth = true;
                        break;
                    case 104:
                        this.screenHeight = x0.readRawVarint32();
                        this.hasScreenHeight = true;
                        break;
                    case 114:
                        int repeatedFieldArrayLength4 = WireFormatNano.getRepeatedFieldArrayLength(x0, 114);
                        int length4 = this.systemSupportedLocale == null ? 0 : this.systemSupportedLocale.length;
                        String[] strArr4 = new String[repeatedFieldArrayLength4 + length4];
                        if (length4 != 0) {
                            System.arraycopy(this.systemSupportedLocale, 0, strArr4, 0, length4);
                        }
                        while (length4 < strArr4.length - 1) {
                            strArr4[length4] = x0.readString();
                            x0.readTag();
                            length4++;
                        }
                        strArr4[length4] = x0.readString();
                        this.systemSupportedLocale = strArr4;
                        break;
                    case 122:
                        int repeatedFieldArrayLength5 = WireFormatNano.getRepeatedFieldArrayLength(x0, 122);
                        int length5 = this.glExtension == null ? 0 : this.glExtension.length;
                        String[] strArr5 = new String[repeatedFieldArrayLength5 + length5];
                        if (length5 != 0) {
                            System.arraycopy(this.glExtension, 0, strArr5, 0, length5);
                        }
                        while (length5 < strArr5.length - 1) {
                            strArr5[length5] = x0.readString();
                            x0.readTag();
                            length5++;
                        }
                        strArr5[length5] = x0.readString();
                        this.glExtension = strArr5;
                        break;
                    case 136:
                        this.maxApkDownloadSizeMb = x0.readRawVarint32();
                        this.hasMaxApkDownloadSizeMb = true;
                        break;
                    case 144:
                        this.smallestScreenWidthDp = x0.readRawVarint32();
                        this.hasSmallestScreenWidthDp = true;
                        break;
                    case 152:
                        this.lowRamDevice = x0.readBool();
                        this.hasLowRamDevice = true;
                        break;
                    case 160:
                        this.totalMemoryBytes = x0.readRawVarint64();
                        this.hasTotalMemoryBytes = true;
                        break;
                    case 168:
                        this.maxNumOfCpuCores = x0.readRawVarint32();
                        this.hasMaxNumOfCpuCores = true;
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

        public DeviceConfigurationProto() {
            this.cachedSize = -1;
        }

        @Override // com.google.protobuf.nano.MessageNano
        public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
            if (this.touchScreen != 0 || this.hasTouchScreen) {
                output.writeInt32(1, this.touchScreen);
            }
            if (this.keyboard != 0 || this.hasKeyboard) {
                output.writeInt32(2, this.keyboard);
            }
            if (this.navigation != 0 || this.hasNavigation) {
                output.writeInt32(3, this.navigation);
            }
            if (this.screenLayout != 0 || this.hasScreenLayout) {
                output.writeInt32(4, this.screenLayout);
            }
            if (this.hasHasHardKeyboard || this.hasHardKeyboard) {
                output.writeBool(5, this.hasHardKeyboard);
            }
            if (this.hasHasFiveWayNavigation || this.hasFiveWayNavigation) {
                output.writeBool(6, this.hasFiveWayNavigation);
            }
            if (this.hasScreenDensity || this.screenDensity != 0) {
                output.writeInt32(7, this.screenDensity);
            }
            if (this.hasGlEsVersion || this.glEsVersion != 0) {
                output.writeInt32(8, this.glEsVersion);
            }
            if (this.systemSharedLibrary != null && this.systemSharedLibrary.length > 0) {
                for (int i = 0; i < this.systemSharedLibrary.length; i++) {
                    String element = this.systemSharedLibrary[i];
                    if (element != null) {
                        output.writeString(9, element);
                    }
                }
            }
            if (this.systemAvailableFeature != null && this.systemAvailableFeature.length > 0) {
                for (int i2 = 0; i2 < this.systemAvailableFeature.length; i2++) {
                    String element2 = this.systemAvailableFeature[i2];
                    if (element2 != null) {
                        output.writeString(10, element2);
                    }
                }
            }
            if (this.nativePlatform != null && this.nativePlatform.length > 0) {
                for (int i3 = 0; i3 < this.nativePlatform.length; i3++) {
                    String element3 = this.nativePlatform[i3];
                    if (element3 != null) {
                        output.writeString(11, element3);
                    }
                }
            }
            if (this.hasScreenWidth || this.screenWidth != 0) {
                output.writeInt32(12, this.screenWidth);
            }
            if (this.hasScreenHeight || this.screenHeight != 0) {
                output.writeInt32(13, this.screenHeight);
            }
            if (this.systemSupportedLocale != null && this.systemSupportedLocale.length > 0) {
                for (int i4 = 0; i4 < this.systemSupportedLocale.length; i4++) {
                    String element4 = this.systemSupportedLocale[i4];
                    if (element4 != null) {
                        output.writeString(14, element4);
                    }
                }
            }
            if (this.glExtension != null && this.glExtension.length > 0) {
                for (int i5 = 0; i5 < this.glExtension.length; i5++) {
                    String element5 = this.glExtension[i5];
                    if (element5 != null) {
                        output.writeString(15, element5);
                    }
                }
            }
            if (this.hasMaxApkDownloadSizeMb || this.maxApkDownloadSizeMb != 50) {
                output.writeInt32(17, this.maxApkDownloadSizeMb);
            }
            if (this.hasSmallestScreenWidthDp || this.smallestScreenWidthDp != 0) {
                output.writeInt32(18, this.smallestScreenWidthDp);
            }
            if (this.hasLowRamDevice || this.lowRamDevice) {
                output.writeBool(19, this.lowRamDevice);
            }
            if (this.hasTotalMemoryBytes || this.totalMemoryBytes != 0) {
                output.writeInt64(20, this.totalMemoryBytes);
            }
            if (this.hasMaxNumOfCpuCores || this.maxNumOfCpuCores != 0) {
                output.writeInt32(21, this.maxNumOfCpuCores);
            }
            super.writeTo(output);
        }

        /* JADX INFO: Access modifiers changed from: protected */
        @Override // com.google.protobuf.nano.MessageNano
        public final int computeSerializedSize() {
            int size = super.computeSerializedSize();
            if (this.touchScreen != 0 || this.hasTouchScreen) {
                size += CodedOutputByteBufferNano.computeInt32Size(1, this.touchScreen);
            }
            if (this.keyboard != 0 || this.hasKeyboard) {
                size += CodedOutputByteBufferNano.computeInt32Size(2, this.keyboard);
            }
            if (this.navigation != 0 || this.hasNavigation) {
                size += CodedOutputByteBufferNano.computeInt32Size(3, this.navigation);
            }
            if (this.screenLayout != 0 || this.hasScreenLayout) {
                size += CodedOutputByteBufferNano.computeInt32Size(4, this.screenLayout);
            }
            if (this.hasHasHardKeyboard || this.hasHardKeyboard) {
                size += CodedOutputByteBufferNano.computeTagSize(5) + 1;
            }
            if (this.hasHasFiveWayNavigation || this.hasFiveWayNavigation) {
                size += CodedOutputByteBufferNano.computeTagSize(6) + 1;
            }
            if (this.hasScreenDensity || this.screenDensity != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(7, this.screenDensity);
            }
            if (this.hasGlEsVersion || this.glEsVersion != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(8, this.glEsVersion);
            }
            if (this.systemSharedLibrary != null && this.systemSharedLibrary.length > 0) {
                int dataCount = 0;
                int dataSize = 0;
                for (int i = 0; i < this.systemSharedLibrary.length; i++) {
                    String element = this.systemSharedLibrary[i];
                    if (element != null) {
                        dataCount++;
                        dataSize += CodedOutputByteBufferNano.computeStringSizeNoTag(element);
                    }
                }
                size = size + dataSize + (dataCount * 1);
            }
            if (this.systemAvailableFeature != null && this.systemAvailableFeature.length > 0) {
                int dataCount2 = 0;
                int dataSize2 = 0;
                for (int i2 = 0; i2 < this.systemAvailableFeature.length; i2++) {
                    String element2 = this.systemAvailableFeature[i2];
                    if (element2 != null) {
                        dataCount2++;
                        dataSize2 += CodedOutputByteBufferNano.computeStringSizeNoTag(element2);
                    }
                }
                size = size + dataSize2 + (dataCount2 * 1);
            }
            if (this.nativePlatform != null && this.nativePlatform.length > 0) {
                int dataCount3 = 0;
                int dataSize3 = 0;
                for (int i3 = 0; i3 < this.nativePlatform.length; i3++) {
                    String element3 = this.nativePlatform[i3];
                    if (element3 != null) {
                        dataCount3++;
                        dataSize3 += CodedOutputByteBufferNano.computeStringSizeNoTag(element3);
                    }
                }
                size = size + dataSize3 + (dataCount3 * 1);
            }
            if (this.hasScreenWidth || this.screenWidth != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(12, this.screenWidth);
            }
            if (this.hasScreenHeight || this.screenHeight != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(13, this.screenHeight);
            }
            if (this.systemSupportedLocale != null && this.systemSupportedLocale.length > 0) {
                int dataCount4 = 0;
                int dataSize4 = 0;
                for (int i4 = 0; i4 < this.systemSupportedLocale.length; i4++) {
                    String element4 = this.systemSupportedLocale[i4];
                    if (element4 != null) {
                        dataCount4++;
                        dataSize4 += CodedOutputByteBufferNano.computeStringSizeNoTag(element4);
                    }
                }
                size = size + dataSize4 + (dataCount4 * 1);
            }
            if (this.glExtension != null && this.glExtension.length > 0) {
                int dataCount5 = 0;
                int dataSize5 = 0;
                for (int i5 = 0; i5 < this.glExtension.length; i5++) {
                    String element5 = this.glExtension[i5];
                    if (element5 != null) {
                        dataCount5++;
                        dataSize5 += CodedOutputByteBufferNano.computeStringSizeNoTag(element5);
                    }
                }
                size = size + dataSize5 + (dataCount5 * 1);
            }
            if (this.hasMaxApkDownloadSizeMb || this.maxApkDownloadSizeMb != 50) {
                size += CodedOutputByteBufferNano.computeInt32Size(17, this.maxApkDownloadSizeMb);
            }
            if (this.hasSmallestScreenWidthDp || this.smallestScreenWidthDp != 0) {
                size += CodedOutputByteBufferNano.computeInt32Size(18, this.smallestScreenWidthDp);
            }
            if (this.hasLowRamDevice || this.lowRamDevice) {
                size += CodedOutputByteBufferNano.computeTagSize(19) + 1;
            }
            if (this.hasTotalMemoryBytes || this.totalMemoryBytes != 0) {
                size += CodedOutputByteBufferNano.computeInt64Size(20, this.totalMemoryBytes);
            }
            if (this.hasMaxNumOfCpuCores || this.maxNumOfCpuCores != 0) {
                return size + CodedOutputByteBufferNano.computeInt32Size(21, this.maxNumOfCpuCores);
            }
            return size;
        }
    }
}
