package com.google.android.gsf.checkin.proto;

import com.google.protobuf.micro.CodedInputStreamMicro;
import com.google.protobuf.micro.CodedOutputStreamMicro;
import com.google.protobuf.micro.MessageMicro;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
/* loaded from: classes.dex */
public final class Checkin$AndroidCheckinRequest extends MessageMicro {
    private boolean hasCheckin;
    private boolean hasDesiredBuild;
    private boolean hasDeviceConfiguration;
    private boolean hasDigest;
    private boolean hasEsn;
    private boolean hasFragment;
    private boolean hasId;
    private boolean hasImei;
    private boolean hasLocale;
    private boolean hasLoggingId;
    private boolean hasMarketCheckin;
    private boolean hasMeid;
    private boolean hasSecurityToken;
    private boolean hasSerialNumber;
    private boolean hasTimeZone;
    private boolean hasUserName;
    private boolean hasUserSerialNumber;
    private boolean hasVersion;
    private String imei_ = "";
    private String meid_ = "";
    private List<String> macAddr_ = Collections.emptyList();
    private List<String> macAddrType_ = Collections.emptyList();
    private String serialNumber_ = "";
    private String esn_ = "";
    private long id_ = 0;
    private long loggingId_ = 0;
    private String digest_ = "";
    private String locale_ = "";
    private Logs$AndroidCheckinProto checkin_ = null;
    private String desiredBuild_ = "";
    private String marketCheckin_ = "";
    private List<String> accountCookie_ = Collections.emptyList();
    private String timeZone_ = "";
    private long securityToken_ = 0;
    private int version_ = 0;
    private List<String> otaCert_ = Collections.emptyList();
    private Config$DeviceConfigurationProto deviceConfiguration_ = null;
    private int fragment_ = 0;
    private String userName_ = "";
    private int userSerialNumber_ = 0;
    private int cachedSize = -1;

    public String getImei() {
        return this.imei_;
    }

    public boolean hasImei() {
        return this.hasImei;
    }

    public Checkin$AndroidCheckinRequest setImei(String value) {
        this.hasImei = true;
        this.imei_ = value;
        return this;
    }

    public String getMeid() {
        return this.meid_;
    }

    public boolean hasMeid() {
        return this.hasMeid;
    }

    public Checkin$AndroidCheckinRequest setMeid(String value) {
        this.hasMeid = true;
        this.meid_ = value;
        return this;
    }

    public List<String> getMacAddrList() {
        return this.macAddr_;
    }

    public Checkin$AndroidCheckinRequest addMacAddr(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.macAddr_.isEmpty()) {
            this.macAddr_ = new ArrayList();
        }
        this.macAddr_.add(value);
        return this;
    }

    public List<String> getMacAddrTypeList() {
        return this.macAddrType_;
    }

    public Checkin$AndroidCheckinRequest addMacAddrType(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.macAddrType_.isEmpty()) {
            this.macAddrType_ = new ArrayList();
        }
        this.macAddrType_.add(value);
        return this;
    }

    public String getSerialNumber() {
        return this.serialNumber_;
    }

    public boolean hasSerialNumber() {
        return this.hasSerialNumber;
    }

    public Checkin$AndroidCheckinRequest setSerialNumber(String value) {
        this.hasSerialNumber = true;
        this.serialNumber_ = value;
        return this;
    }

    public String getEsn() {
        return this.esn_;
    }

    public boolean hasEsn() {
        return this.hasEsn;
    }

    public Checkin$AndroidCheckinRequest setEsn(String value) {
        this.hasEsn = true;
        this.esn_ = value;
        return this;
    }

    public long getId() {
        return this.id_;
    }

    public boolean hasId() {
        return this.hasId;
    }

    public Checkin$AndroidCheckinRequest setId(long value) {
        this.hasId = true;
        this.id_ = value;
        return this;
    }

    public long getLoggingId() {
        return this.loggingId_;
    }

    public boolean hasLoggingId() {
        return this.hasLoggingId;
    }

    public Checkin$AndroidCheckinRequest setLoggingId(long value) {
        this.hasLoggingId = true;
        this.loggingId_ = value;
        return this;
    }

    public String getDigest() {
        return this.digest_;
    }

    public boolean hasDigest() {
        return this.hasDigest;
    }

    public Checkin$AndroidCheckinRequest setDigest(String value) {
        this.hasDigest = true;
        this.digest_ = value;
        return this;
    }

    public String getLocale() {
        return this.locale_;
    }

    public boolean hasLocale() {
        return this.hasLocale;
    }

    public Checkin$AndroidCheckinRequest setLocale(String value) {
        this.hasLocale = true;
        this.locale_ = value;
        return this;
    }

    public boolean hasCheckin() {
        return this.hasCheckin;
    }

    public Logs$AndroidCheckinProto getCheckin() {
        return this.checkin_;
    }

    public Checkin$AndroidCheckinRequest setCheckin(Logs$AndroidCheckinProto value) {
        if (value == null) {
            throw new NullPointerException();
        }
        this.hasCheckin = true;
        this.checkin_ = value;
        return this;
    }

    public String getDesiredBuild() {
        return this.desiredBuild_;
    }

    public boolean hasDesiredBuild() {
        return this.hasDesiredBuild;
    }

    public Checkin$AndroidCheckinRequest setDesiredBuild(String value) {
        this.hasDesiredBuild = true;
        this.desiredBuild_ = value;
        return this;
    }

    public String getMarketCheckin() {
        return this.marketCheckin_;
    }

    public boolean hasMarketCheckin() {
        return this.hasMarketCheckin;
    }

    public Checkin$AndroidCheckinRequest setMarketCheckin(String value) {
        this.hasMarketCheckin = true;
        this.marketCheckin_ = value;
        return this;
    }

    public List<String> getAccountCookieList() {
        return this.accountCookie_;
    }

    public Checkin$AndroidCheckinRequest addAccountCookie(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.accountCookie_.isEmpty()) {
            this.accountCookie_ = new ArrayList();
        }
        this.accountCookie_.add(value);
        return this;
    }

    public String getTimeZone() {
        return this.timeZone_;
    }

    public boolean hasTimeZone() {
        return this.hasTimeZone;
    }

    public Checkin$AndroidCheckinRequest setTimeZone(String value) {
        this.hasTimeZone = true;
        this.timeZone_ = value;
        return this;
    }

    public long getSecurityToken() {
        return this.securityToken_;
    }

    public boolean hasSecurityToken() {
        return this.hasSecurityToken;
    }

    public Checkin$AndroidCheckinRequest setSecurityToken(long value) {
        this.hasSecurityToken = true;
        this.securityToken_ = value;
        return this;
    }

    public int getVersion() {
        return this.version_;
    }

    public boolean hasVersion() {
        return this.hasVersion;
    }

    public Checkin$AndroidCheckinRequest setVersion(int value) {
        this.hasVersion = true;
        this.version_ = value;
        return this;
    }

    public List<String> getOtaCertList() {
        return this.otaCert_;
    }

    public int getOtaCertCount() {
        return this.otaCert_.size();
    }

    public Checkin$AndroidCheckinRequest addOtaCert(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.otaCert_.isEmpty()) {
            this.otaCert_ = new ArrayList();
        }
        this.otaCert_.add(value);
        return this;
    }

    public boolean hasDeviceConfiguration() {
        return this.hasDeviceConfiguration;
    }

    public Config$DeviceConfigurationProto getDeviceConfiguration() {
        return this.deviceConfiguration_;
    }

    public Checkin$AndroidCheckinRequest setDeviceConfiguration(Config$DeviceConfigurationProto value) {
        if (value == null) {
            throw new NullPointerException();
        }
        this.hasDeviceConfiguration = true;
        this.deviceConfiguration_ = value;
        return this;
    }

    public int getFragment() {
        return this.fragment_;
    }

    public boolean hasFragment() {
        return this.hasFragment;
    }

    public Checkin$AndroidCheckinRequest setFragment(int value) {
        this.hasFragment = true;
        this.fragment_ = value;
        return this;
    }

    public String getUserName() {
        return this.userName_;
    }

    public boolean hasUserName() {
        return this.hasUserName;
    }

    public Checkin$AndroidCheckinRequest setUserName(String value) {
        this.hasUserName = true;
        this.userName_ = value;
        return this;
    }

    public int getUserSerialNumber() {
        return this.userSerialNumber_;
    }

    public boolean hasUserSerialNumber() {
        return this.hasUserSerialNumber;
    }

    public Checkin$AndroidCheckinRequest setUserSerialNumber(int value) {
        this.hasUserSerialNumber = true;
        this.userSerialNumber_ = value;
        return this;
    }

    public final boolean isInitialized() {
        if (this.hasCheckin && getCheckin().isInitialized()) {
            return !hasDeviceConfiguration() || getDeviceConfiguration().isInitialized();
        }
        return false;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public void writeTo(CodedOutputStreamMicro output) throws IOException {
        if (hasImei()) {
            output.writeString(1, getImei());
        }
        if (hasId()) {
            output.writeInt64(2, getId());
        }
        if (hasDigest()) {
            output.writeString(3, getDigest());
        }
        if (hasCheckin()) {
            output.writeMessage(4, getCheckin());
        }
        if (hasDesiredBuild()) {
            output.writeString(5, getDesiredBuild());
        }
        if (hasLocale()) {
            output.writeString(6, getLocale());
        }
        if (hasLoggingId()) {
            output.writeInt64(7, getLoggingId());
        }
        if (hasMarketCheckin()) {
            output.writeString(8, getMarketCheckin());
        }
        for (String element : getMacAddrList()) {
            output.writeString(9, element);
        }
        if (hasMeid()) {
            output.writeString(10, getMeid());
        }
        for (String element2 : getAccountCookieList()) {
            output.writeString(11, element2);
        }
        if (hasTimeZone()) {
            output.writeString(12, getTimeZone());
        }
        if (hasSecurityToken()) {
            output.writeFixed64(13, getSecurityToken());
        }
        if (hasVersion()) {
            output.writeInt32(14, getVersion());
        }
        for (String element3 : getOtaCertList()) {
            output.writeString(15, element3);
        }
        if (hasSerialNumber()) {
            output.writeString(16, getSerialNumber());
        }
        if (hasEsn()) {
            output.writeString(17, getEsn());
        }
        if (hasDeviceConfiguration()) {
            output.writeMessage(18, getDeviceConfiguration());
        }
        for (String element4 : getMacAddrTypeList()) {
            output.writeString(19, element4);
        }
        if (hasFragment()) {
            output.writeInt32(20, getFragment());
        }
        if (hasUserName()) {
            output.writeString(21, getUserName());
        }
        if (!hasUserSerialNumber()) {
            return;
        }
        output.writeInt32(22, getUserSerialNumber());
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
        if (hasImei()) {
            size = CodedOutputStreamMicro.computeStringSize(1, getImei()) + 0;
        }
        if (hasId()) {
            size += CodedOutputStreamMicro.computeInt64Size(2, getId());
        }
        if (hasDigest()) {
            size += CodedOutputStreamMicro.computeStringSize(3, getDigest());
        }
        if (hasCheckin()) {
            size += CodedOutputStreamMicro.computeMessageSize(4, getCheckin());
        }
        if (hasDesiredBuild()) {
            size += CodedOutputStreamMicro.computeStringSize(5, getDesiredBuild());
        }
        if (hasLocale()) {
            size += CodedOutputStreamMicro.computeStringSize(6, getLocale());
        }
        if (hasLoggingId()) {
            size += CodedOutputStreamMicro.computeInt64Size(7, getLoggingId());
        }
        if (hasMarketCheckin()) {
            size += CodedOutputStreamMicro.computeStringSize(8, getMarketCheckin());
        }
        int dataSize = 0;
        for (String element : getMacAddrList()) {
            dataSize += CodedOutputStreamMicro.computeStringSizeNoTag(element);
        }
        int size2 = size + dataSize + (getMacAddrList().size() * 1);
        if (hasMeid()) {
            size2 += CodedOutputStreamMicro.computeStringSize(10, getMeid());
        }
        int dataSize2 = 0;
        for (String element2 : getAccountCookieList()) {
            dataSize2 += CodedOutputStreamMicro.computeStringSizeNoTag(element2);
        }
        int size3 = size2 + dataSize2 + (getAccountCookieList().size() * 1);
        if (hasTimeZone()) {
            size3 += CodedOutputStreamMicro.computeStringSize(12, getTimeZone());
        }
        if (hasSecurityToken()) {
            size3 += CodedOutputStreamMicro.computeFixed64Size(13, getSecurityToken());
        }
        if (hasVersion()) {
            size3 += CodedOutputStreamMicro.computeInt32Size(14, getVersion());
        }
        int dataSize3 = 0;
        for (String element3 : getOtaCertList()) {
            dataSize3 += CodedOutputStreamMicro.computeStringSizeNoTag(element3);
        }
        int size4 = size3 + dataSize3 + (getOtaCertList().size() * 1);
        if (hasSerialNumber()) {
            size4 += CodedOutputStreamMicro.computeStringSize(16, getSerialNumber());
        }
        if (hasEsn()) {
            size4 += CodedOutputStreamMicro.computeStringSize(17, getEsn());
        }
        if (hasDeviceConfiguration()) {
            size4 += CodedOutputStreamMicro.computeMessageSize(18, getDeviceConfiguration());
        }
        int dataSize4 = 0;
        for (String element4 : getMacAddrTypeList()) {
            dataSize4 += CodedOutputStreamMicro.computeStringSizeNoTag(element4);
        }
        int size5 = size4 + dataSize4 + (getMacAddrTypeList().size() * 2);
        if (hasFragment()) {
            size5 += CodedOutputStreamMicro.computeInt32Size(20, getFragment());
        }
        if (hasUserName()) {
            size5 += CodedOutputStreamMicro.computeStringSize(21, getUserName());
        }
        if (hasUserSerialNumber()) {
            size5 += CodedOutputStreamMicro.computeInt32Size(22, getUserSerialNumber());
        }
        this.cachedSize = size5;
        return size5;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public Checkin$AndroidCheckinRequest mergeFrom(CodedInputStreamMicro input) throws IOException {
        while (true) {
            int tag = input.readTag();
            switch (tag) {
                case 0:
                    return this;
                case 10:
                    setImei(input.readString());
                    break;
                case 16:
                    setId(input.readInt64());
                    break;
                case 26:
                    setDigest(input.readString());
                    break;
                case 34:
                    Logs$AndroidCheckinProto value = new Logs$AndroidCheckinProto();
                    input.readMessage(value);
                    setCheckin(value);
                    break;
                case 42:
                    setDesiredBuild(input.readString());
                    break;
                case 50:
                    setLocale(input.readString());
                    break;
                case 56:
                    setLoggingId(input.readInt64());
                    break;
                case 66:
                    setMarketCheckin(input.readString());
                    break;
                case 74:
                    addMacAddr(input.readString());
                    break;
                case 82:
                    setMeid(input.readString());
                    break;
                case 90:
                    addAccountCookie(input.readString());
                    break;
                case 98:
                    setTimeZone(input.readString());
                    break;
                case 105:
                    setSecurityToken(input.readFixed64());
                    break;
                case 112:
                    setVersion(input.readInt32());
                    break;
                case 122:
                    addOtaCert(input.readString());
                    break;
                case 130:
                    setSerialNumber(input.readString());
                    break;
                case 138:
                    setEsn(input.readString());
                    break;
                case 146:
                    Config$DeviceConfigurationProto value2 = new Config$DeviceConfigurationProto();
                    input.readMessage(value2);
                    setDeviceConfiguration(value2);
                    break;
                case 154:
                    addMacAddrType(input.readString());
                    break;
                case 160:
                    setFragment(input.readInt32());
                    break;
                case 170:
                    setUserName(input.readString());
                    break;
                case 176:
                    setUserSerialNumber(input.readInt32());
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
