package com.google.android.finsky.protos;

import com.google.protobuf.nano.CodedInputByteBufferNano;
import com.google.protobuf.nano.CodedOutputByteBufferNano;
import com.google.protobuf.nano.MessageNano;
import com.google.protobuf.nano.WireFormatNano;
import java.io.IOException;
/* loaded from: classes.dex */
public final class InstallDetails extends MessageNano {
    public int installLocation = 0;
    public boolean hasInstallLocation = false;
    public long size = 0;
    public boolean hasSize = false;
    public Dependency[] dependency = Dependency.emptyArray();
    public int targetSdkVersion = 0;
    public boolean hasTargetSdkVersion = false;

    @Override // com.google.protobuf.nano.MessageNano
    public final /* bridge */ /* synthetic */ MessageNano mergeFrom(CodedInputByteBufferNano x0) throws IOException {
        int length;
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
                            this.installLocation = readRawVarint32;
                            this.hasInstallLocation = true;
                            continue;
                    }
                case 16:
                    this.size = x0.readRawVarint64();
                    this.hasSize = true;
                    break;
                case 26:
                    int repeatedFieldArrayLength = WireFormatNano.getRepeatedFieldArrayLength(x0, 26);
                    if (this.dependency == null) {
                        length = 0;
                    } else {
                        length = this.dependency.length;
                    }
                    Dependency[] dependencyArr = new Dependency[repeatedFieldArrayLength + length];
                    if (length != 0) {
                        System.arraycopy(this.dependency, 0, dependencyArr, 0, length);
                    }
                    while (length < dependencyArr.length - 1) {
                        dependencyArr[length] = new Dependency();
                        x0.readMessage(dependencyArr[length]);
                        x0.readTag();
                        length++;
                    }
                    dependencyArr[length] = new Dependency();
                    x0.readMessage(dependencyArr[length]);
                    this.dependency = dependencyArr;
                    break;
                case 32:
                    this.targetSdkVersion = x0.readRawVarint32();
                    this.hasTargetSdkVersion = true;
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

    public InstallDetails() {
        this.cachedSize = -1;
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final void writeTo(CodedOutputByteBufferNano output) throws IOException {
        if (this.installLocation != 0 || this.hasInstallLocation) {
            output.writeInt32(1, this.installLocation);
        }
        if (this.hasSize || this.size != 0) {
            output.writeInt64(2, this.size);
        }
        if (this.dependency != null && this.dependency.length > 0) {
            for (int i = 0; i < this.dependency.length; i++) {
                Dependency element = this.dependency[i];
                if (element != null) {
                    output.writeMessage(3, element);
                }
            }
        }
        if (this.hasTargetSdkVersion || this.targetSdkVersion != 0) {
            output.writeInt32(4, this.targetSdkVersion);
        }
        super.writeTo(output);
    }

    @Override // com.google.protobuf.nano.MessageNano
    public final int computeSerializedSize() {
        int size = super.computeSerializedSize();
        if (this.installLocation != 0 || this.hasInstallLocation) {
            size += CodedOutputByteBufferNano.computeInt32Size(1, this.installLocation);
        }
        if (this.hasSize || this.size != 0) {
            size += CodedOutputByteBufferNano.computeInt64Size(2, this.size);
        }
        if (this.dependency != null && this.dependency.length > 0) {
            for (int i = 0; i < this.dependency.length; i++) {
                Dependency element = this.dependency[i];
                if (element != null) {
                    size += CodedOutputByteBufferNano.computeMessageSize(3, element);
                }
            }
        }
        if (this.hasTargetSdkVersion || this.targetSdkVersion != 0) {
            return size + CodedOutputByteBufferNano.computeInt32Size(4, this.targetSdkVersion);
        }
        return size;
    }
}
