package com.google.android.gsf.checkin.proto;

import com.google.protobuf.micro.CodedInputStreamMicro;
import com.google.protobuf.micro.CodedOutputStreamMicro;
import com.google.protobuf.micro.MessageMicro;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
/* loaded from: classes.dex */
public final class Logs$AndroidCheckinProto extends MessageMicro {
    private boolean hasBuild;
    private boolean hasCellOperator;
    private boolean hasLastCheckinMsec;
    private boolean hasRoaming;
    private boolean hasSimOperator;
    private boolean hasUserNumber;
    private Logs$AndroidBuildProto build_ = null;
    private long lastCheckinMsec_ = 0;
    private List<String> requestedGroup_ = Collections.emptyList();
    private List<Logs$AndroidEventProto> event_ = Collections.emptyList();
    private List<Logs$AndroidStatisticProto> stat_ = Collections.emptyList();
    private String cellOperator_ = "";
    private String simOperator_ = "";
    private String roaming_ = "";
    private int userNumber_ = 0;
    private int cachedSize = -1;

    public boolean hasBuild() {
        return this.hasBuild;
    }

    public Logs$AndroidBuildProto getBuild() {
        return this.build_;
    }

    public Logs$AndroidCheckinProto setBuild(Logs$AndroidBuildProto value) {
        if (value == null) {
            throw new NullPointerException();
        }
        this.hasBuild = true;
        this.build_ = value;
        return this;
    }

    public long getLastCheckinMsec() {
        return this.lastCheckinMsec_;
    }

    public boolean hasLastCheckinMsec() {
        return this.hasLastCheckinMsec;
    }

    public Logs$AndroidCheckinProto setLastCheckinMsec(long value) {
        this.hasLastCheckinMsec = true;
        this.lastCheckinMsec_ = value;
        return this;
    }

    public List<String> getRequestedGroupList() {
        return this.requestedGroup_;
    }

    public Logs$AndroidCheckinProto addRequestedGroup(String value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.requestedGroup_.isEmpty()) {
            this.requestedGroup_ = new ArrayList();
        }
        this.requestedGroup_.add(value);
        return this;
    }

    public List<Logs$AndroidEventProto> getEventList() {
        return this.event_;
    }

    public int getEventCount() {
        return this.event_.size();
    }

    public Logs$AndroidEventProto getEvent(int index) {
        return this.event_.get(index);
    }

    public Logs$AndroidCheckinProto addEvent(Logs$AndroidEventProto value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.event_.isEmpty()) {
            this.event_ = new ArrayList();
        }
        this.event_.add(value);
        return this;
    }

    public List<Logs$AndroidStatisticProto> getStatList() {
        return this.stat_;
    }

    public Logs$AndroidCheckinProto addStat(Logs$AndroidStatisticProto value) {
        if (value == null) {
            throw new NullPointerException();
        }
        if (this.stat_.isEmpty()) {
            this.stat_ = new ArrayList();
        }
        this.stat_.add(value);
        return this;
    }

    public String getCellOperator() {
        return this.cellOperator_;
    }

    public boolean hasCellOperator() {
        return this.hasCellOperator;
    }

    public Logs$AndroidCheckinProto setCellOperator(String value) {
        this.hasCellOperator = true;
        this.cellOperator_ = value;
        return this;
    }

    public String getSimOperator() {
        return this.simOperator_;
    }

    public boolean hasSimOperator() {
        return this.hasSimOperator;
    }

    public Logs$AndroidCheckinProto setSimOperator(String value) {
        this.hasSimOperator = true;
        this.simOperator_ = value;
        return this;
    }

    public String getRoaming() {
        return this.roaming_;
    }

    public boolean hasRoaming() {
        return this.hasRoaming;
    }

    public Logs$AndroidCheckinProto setRoaming(String value) {
        this.hasRoaming = true;
        this.roaming_ = value;
        return this;
    }

    public int getUserNumber() {
        return this.userNumber_;
    }

    public boolean hasUserNumber() {
        return this.hasUserNumber;
    }

    public Logs$AndroidCheckinProto setUserNumber(int value) {
        this.hasUserNumber = true;
        this.userNumber_ = value;
        return this;
    }

    public final boolean isInitialized() {
        for (Logs$AndroidEventProto element : getEventList()) {
            if (!element.isInitialized()) {
                return false;
            }
        }
        for (Logs$AndroidStatisticProto element2 : getStatList()) {
            if (!element2.isInitialized()) {
                return false;
            }
        }
        return true;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public void writeTo(CodedOutputStreamMicro output) throws IOException {
        if (hasBuild()) {
            output.writeMessage(1, getBuild());
        }
        if (hasLastCheckinMsec()) {
            output.writeInt64(2, getLastCheckinMsec());
        }
        for (Logs$AndroidEventProto element : getEventList()) {
            output.writeMessage(3, element);
        }
        for (Logs$AndroidStatisticProto element2 : getStatList()) {
            output.writeMessage(4, element2);
        }
        for (String element3 : getRequestedGroupList()) {
            output.writeString(5, element3);
        }
        if (hasCellOperator()) {
            output.writeString(6, getCellOperator());
        }
        if (hasSimOperator()) {
            output.writeString(7, getSimOperator());
        }
        if (hasRoaming()) {
            output.writeString(8, getRoaming());
        }
        if (!hasUserNumber()) {
            return;
        }
        output.writeInt32(9, getUserNumber());
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
        if (hasBuild()) {
            size = CodedOutputStreamMicro.computeMessageSize(1, getBuild()) + 0;
        }
        if (hasLastCheckinMsec()) {
            size += CodedOutputStreamMicro.computeInt64Size(2, getLastCheckinMsec());
        }
        for (Logs$AndroidEventProto element : getEventList()) {
            size += CodedOutputStreamMicro.computeMessageSize(3, element);
        }
        for (Logs$AndroidStatisticProto element2 : getStatList()) {
            size += CodedOutputStreamMicro.computeMessageSize(4, element2);
        }
        int dataSize = 0;
        for (String element3 : getRequestedGroupList()) {
            dataSize += CodedOutputStreamMicro.computeStringSizeNoTag(element3);
        }
        int size2 = size + dataSize + (getRequestedGroupList().size() * 1);
        if (hasCellOperator()) {
            size2 += CodedOutputStreamMicro.computeStringSize(6, getCellOperator());
        }
        if (hasSimOperator()) {
            size2 += CodedOutputStreamMicro.computeStringSize(7, getSimOperator());
        }
        if (hasRoaming()) {
            size2 += CodedOutputStreamMicro.computeStringSize(8, getRoaming());
        }
        if (hasUserNumber()) {
            size2 += CodedOutputStreamMicro.computeInt32Size(9, getUserNumber());
        }
        this.cachedSize = size2;
        return size2;
    }

    @Override // com.google.protobuf.micro.MessageMicro
    public Logs$AndroidCheckinProto mergeFrom(CodedInputStreamMicro input) throws IOException {
        while (true) {
            int tag = input.readTag();
            switch (tag) {
                case 0:
                    return this;
                case 10:
                    Logs$AndroidBuildProto value = new Logs$AndroidBuildProto();
                    input.readMessage(value);
                    setBuild(value);
                    break;
                case 16:
                    setLastCheckinMsec(input.readInt64());
                    break;
                case 26:
                    Logs$AndroidEventProto value2 = new Logs$AndroidEventProto();
                    input.readMessage(value2);
                    addEvent(value2);
                    break;
                case 34:
                    Logs$AndroidStatisticProto value3 = new Logs$AndroidStatisticProto();
                    input.readMessage(value3);
                    addStat(value3);
                    break;
                case 42:
                    addRequestedGroup(input.readString());
                    break;
                case 50:
                    setCellOperator(input.readString());
                    break;
                case 58:
                    setSimOperator(input.readString());
                    break;
                case 66:
                    setRoaming(input.readString());
                    break;
                case 72:
                    setUserNumber(input.readInt32());
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
