package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class ServerCommands extends MessageNano {
    public boolean clearCache = false;
    public boolean hasClearCache = false;
    public String displayErrorMessage = "";
    public boolean hasDisplayErrorMessage = false;
    public String logErrorStacktrace = "";
    public boolean hasLogErrorStacktrace = false;
    public UserSettingDirtyData[] userSettingDirtyData = UserSettingDirtyData.emptyArray();

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
        while (true) {
            int readTag = x0.readTag();
            switch (readTag) {
                case 0:
                    break;
                case 8:
                    this.clearCache = x0.readBool();
                    this.hasClearCache = true;
                    break;
                case 18:
                    this.displayErrorMessage = x0.readString();
                    this.hasDisplayErrorMessage = true;
                    break;
                case 26:
                    this.logErrorStacktrace = x0.readString();
                    this.hasLogErrorStacktrace = true;
                    break;
                case 34:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 34);
                    if (this.userSettingDirtyData == null) {
                        length = 0;
                    } else {
                        length = this.userSettingDirtyData.length;
                    }
                    UserSettingDirtyData[] userSettingDirtyDataArr = new UserSettingDirtyData[repeatedFieldArrayLength + length];
                    if (length != 0) {
                        System.arraycopy(this.userSettingDirtyData, 0, userSettingDirtyDataArr, 0, length);
                    }
                    while (length < userSettingDirtyDataArr.length - 1) {
                        userSettingDirtyDataArr[length] = new UserSettingDirtyData();
                        x0.readMessage(userSettingDirtyDataArr[length]);
                        x0.readTag();
                        length++;
                    }
                    userSettingDirtyDataArr[length] = new UserSettingDirtyData();
                    x0.readMessage(userSettingDirtyDataArr[length]);
                    this.userSettingDirtyData = userSettingDirtyDataArr;
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

    public ServerCommands() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.hasClearCache || this.clearCache) {
            output.writeBool(1, this.clearCache);
        }
        if (this.hasDisplayErrorMessage || !this.displayErrorMessage.equals("")) {
            output.writeString(2, this.displayErrorMessage);
        }
        if (this.hasLogErrorStacktrace || !this.logErrorStacktrace.equals("")) {
            output.writeString(3, this.logErrorStacktrace);
        }
        if (this.userSettingDirtyData != null && this.userSettingDirtyData.length > 0) {
            for (int i = 0; i < this.userSettingDirtyData.length; i++) {
                UserSettingDirtyData element = this.userSettingDirtyData[i];
                if (element != null) {
                    output.writeMessage(4, element);
                }
            }
        }
        super.writeTo(output);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.hasClearCache || this.clearCache) {
            size += CodedOutputByteBufferNano.computeTagSize(1) + 1;
        }
        if (this.hasDisplayErrorMessage || !this.displayErrorMessage.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(2, this.displayErrorMessage);
        }
        if (this.hasLogErrorStacktrace || !this.logErrorStacktrace.equals("")) {
            size += CodedOutputByteBufferNano.computeStringSize(3, this.logErrorStacktrace);
        }
        if (this.userSettingDirtyData != null && this.userSettingDirtyData.length > 0) {
            for (int i = 0; i < this.userSettingDirtyData.length; i++) {
                UserSettingDirtyData element = this.userSettingDirtyData[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(4, element);
                }
            }
        }
        return size;
    }
}
