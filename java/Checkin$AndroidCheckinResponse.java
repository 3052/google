package com.google.android.gsf.checkin.proto;

import com.google.protobuf.micro.CodedInputStreamMicro;
import com.google.protobuf.micro.CodedOutputStreamMicro;
import com.google.protobuf.micro.MessageMicro;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
/* loaded from: classes.dex */
public final class Checkin$AndroidCheckinResponse extends MessageMicro {
    private boolean hasAndroidId;
    private boolean hasDigest;
    private boolean hasMarketOk;
    private boolean hasSecurityToken;
    private boolean hasSettingsDiff;
    private boolean hasStatsOk;
    private boolean hasTimeMsec;
    private boolean statsOk_ = false;
    private long timeMsec_ = 0;
    private List<Logs$AndroidIntentProto> intent_ = Collections.emptyList();
    private String digest_ = "";
    private boolean settingsDiff_ = false;
    private List<String> deleteSetting_ = Collections.emptyList();
    private List<Checkin$GservicesSetting> setting_ = Collections.emptyList();
    private boolean marketOk_ = false;
    private long androidId_ = 0;
    private long securityToken_ = 0;
    private int cachedSize = -1;

    public boolean getStatsOk() {
        return this.statsOk_;
    }

    public boolean hasStatsOk() {
        return this.hasStatsOk;
    }

    public Checkin$AndroidCheckinResponse setStatsOk(boolean value) {
        this.hasStatsOk = true;
        this.statsOk_ = value;
        return this;
    }

    public long getTimeMsec() {
        return this.timeMsec_;
    }

    public boolean hasTimeMsec() {
        return this.hasTimeMsec;
    }

    public Checkin$AndroidCheckinResponse setTimeMsec(long value) {
        this.hasTimeMsec = true;
        this.timeMsec_ = value;
        return this;
    }

    public List<Logs$AndroidIntentProto> getIntentList() {
        return this.intent_;
    }

    public int getIntentCount() {
        return this.intent_.size();
    }

    public Logs$AndroidIntentProto getIntent(int index) {
        return this.intent_.get(index);
    }

    public Checkin$AndroidCheckinResponse addIntent(Logs$AndroidIntentProto value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.intent_.isEmpty()) {
            this.intent_ = new ArrayList();
        }
        this.intent_.add(value);
        return this;
    }

    public String getDigest() {
        return this.digest_;
    }

    public boolean hasDigest() {
        return this.hasDigest;
    }

    public Checkin$AndroidCheckinResponse setDigest(String value) {
        this.hasDigest = true;
        this.digest_ = value;
        return this;
    }

    public boolean getSettingsDiff() {
        return this.settingsDiff_;
    }

    public boolean hasSettingsDiff() {
        return this.hasSettingsDiff;
    }

    public Checkin$AndroidCheckinResponse setSettingsDiff(boolean value) {
        this.hasSettingsDiff = true;
        this.settingsDiff_ = value;
        return this;
    }

    public List<String> getDeleteSettingList() {
        return this.deleteSetting_;
    }

    public int getDeleteSettingCount() {
        return this.deleteSetting_.size();
    }

    public String getDeleteSetting(int index) {
        return this.deleteSetting_.get(index);
    }

    public Checkin$AndroidCheckinResponse addDeleteSetting(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.deleteSetting_.isEmpty()) {
            this.deleteSetting_ = new ArrayList();
        }
        this.deleteSetting_.add(value);
        return this;
    }

    public List<Checkin$GservicesSetting> getSettingList() {
        return this.setting_;
    }

    public int getSettingCount() {
        return this.setting_.size();
    }

    public Checkin$GservicesSetting getSetting(int index) {
        return this.setting_.get(index);
    }

    public Checkin$AndroidCheckinResponse addSetting(Checkin$GservicesSetting value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.setting_.isEmpty()) {
            this.setting_ = new ArrayList();
        }
        this.setting_.add(value);
        return this;
    }

    public boolean getMarketOk() {
        return this.marketOk_;
    }

    public boolean hasMarketOk() {
        return this.hasMarketOk;
    }

    public Checkin$AndroidCheckinResponse setMarketOk(boolean value) {
        this.hasMarketOk = true;
        this.marketOk_ = value;
        return this;
    }

    public long getAndroidId() {
        return this.androidId_;
    }

    public boolean hasAndroidId() {
        return this.hasAndroidId;
    }

    public Checkin$AndroidCheckinResponse setAndroidId(long value) {
        this.hasAndroidId = true;
        this.androidId_ = value;
        return this;
    }

    public long getSecurityToken() {
        return this.securityToken_;
    }

    public boolean hasSecurityToken() {
        return this.hasSecurityToken;
    }

    public Checkin$AndroidCheckinResponse setSecurityToken(long value) {
        this.hasSecurityToken = true;
        this.securityToken_ = value;
        return this;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public void writeTo(CodedOutputStreamMicro output) throws IOException {
        if (hasStatsOk()) {
            output.writeBool(1, getStatsOk());
        }
        for (Logs$AndroidIntentProto element : getIntentList()) {
            output.writeMessage(2, element);
        }
        if (hasTimeMsec()) {
            output.writeInt64(3, getTimeMsec());
        }
        if (hasDigest()) {
            output.writeString(4, getDigest());
        }
        for (Checkin$GservicesSetting element2 : getSettingList()) {
            output.writeMessage(5, element2);
        }
        if (hasMarketOk()) {
            output.writeBool(6, getMarketOk());
        }
        if (hasAndroidId()) {
            output.writeFixed64(7, getAndroidId());
        }
        if (hasSecurityToken()) {
            output.writeFixed64(8, getSecurityToken());
        }
        if (hasSettingsDiff()) {
            output.writeBool(9, getSettingsDiff());
        }
        for (String element3 : getDeleteSettingList()) {
            output.writeString(10, element3);
        }
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public int getCachedSize() {
        if (this.cachedSize < 0) {
            getSerializedSize();
        }
        return this.cachedSize;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public int getSerializedSize() {
        int size = 0;
        if (hasStatsOk()) {
            size = CodedOutputStreamMicro.computeBoolSize(1, getStatsOk()) + 0;
        }
        for (Logs$AndroidIntentProto element : getIntentList()) {
            size += CodedOutputStreamMicro.computeMessageSize(2, element);
        }
        if (hasTimeMsec()) {
            size += CodedOutputStreamMicro.computeInt64Size(3, getTimeMsec());
        }
        if (hasDigest()) {
            size += CodedOutputStreamMicro.computeStringSize(4, getDigest());
        }
        for (Checkin$GservicesSetting element2 : getSettingList()) {
            size += CodedOutputStreamMicro.computeMessageSize(5, element2);
        }
        if (hasMarketOk()) {
            size += CodedOutputStreamMicro.computeBoolSize(6, getMarketOk());
        }
        if (hasAndroidId()) {
            size += CodedOutputStreamMicro.computeFixed64Size(7, getAndroidId());
        }
        if (hasSecurityToken()) {
            size += CodedOutputStreamMicro.computeFixed64Size(8, getSecurityToken());
        }
        if (hasSettingsDiff()) {
            size += CodedOutputStreamMicro.computeBoolSize(9, getSettingsDiff());
        }
        int dataSize = 0;
        for (String element3 : getDeleteSettingList()) {
            dataSize += CodedOutputStreamMicro.computeStringSizeNoTag(element3);
        }
        int size2 = size + dataSize + (getDeleteSettingList().size() * 1);
        this.cachedSize = size2;
        return size2;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public Checkin$AndroidCheckinResponse mergeFrom(CodedInputStreamMicro input) throws IOException {
        while (true) {
            int tag = input.readTag();
            switch (tag) {
                case 0:
                    return this;
                case 8:
                    setStatsOk(input.readBool());
                    break;
                case 18:
                    Logs$AndroidIntentProto value = new Logs$AndroidIntentProto();
                    input.readMessage(value);
                    addIntent(value);
                    break;
                case 24:
                    setTimeMsec(input.readInt64());
                    break;
                case 34:
                    setDigest(input.readString());
                    break;
                case 42:
                    Checkin$GservicesSetting value2 = new Checkin$GservicesSetting();
                    input.readMessage(value2);
                    addSetting(value2);
                    break;
                case 48:
                    setMarketOk(input.readBool());
                    break;
                case 57:
                    setAndroidId(input.readFixed64());
                    break;
                case 65:
                    setSecurityToken(input.readFixed64());
                    break;
                case 72:
                    setSettingsDiff(input.readBool());
                    break;
                case 82:
                    addDeleteSetting(input.readString());
                    break;
                default:
                    if (parseUnknownField(input, tag)) {
                        break;
                    } else {
                        return this;
                    }
            }
        }
    }
}
