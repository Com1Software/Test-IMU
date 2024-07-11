#!/usr/bin/env python3
import math
import re
import serial
import struct

AngInit = b"\xff\xaa\x52"
AccCalib = b"\xff\xaa\x67"
declination = -0.00669
pi          = 3.14159265359
feature = b"UQ(.{6,6}).{3,3}UR(.{6,6}).{3,3}US(.{6,6}).{3,3}"
fmt_B, fmt_h = "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB", "<hhh"

s_x, s_y, s_z = 0.0, 0.0, 0.0

ser = serial.Serial('/dev/ttyUSB0',baudrate=115200)
if ser.isOpen():
    print("success")
else:
    print("failed")

ser.write(AngInit)
ser.write(AccCalib)

while True:
   imu_msg = ser.read(size=65)
   result = re.search(feature, imu_msg)
   if result:
        frame = struct.unpack(fmt_B, result.group())
        hex_string = "".join("%02x " % b for b in frame)
        sum_Q, sum_R, sum_S = 0, 0, 0
        for i in range(0, 10):
            sum_Q, sum_R, sum_S = sum_Q+frame[i], sum_R+frame[i+11], sum_S+frame[i+22]
            sum_Q, sum_R, sum_S = sum_Q&0x000000ff, sum_R&0x000000ff, sum_S&0x000000ff

        if (sum_Q==frame[10]) and (sum_R==frame[21]) and (sum_S==frame[32]):
            af, wf, ef = struct.unpack(fmt_h, result.group(1)), struct.unpack(fmt_h, result.group(2)), struct.unpack(fmt_h, result.group(3))
          
            af_l, wf_l, ef_l = [], [], []
            for i in range(0, 3):
                af_l.append(round(af[i]/32768.0*16,2)*9.8), wf_l.append(wf[i]/32768.0*2000), ef_l.append(round(ef[i]/32768.0*180,2))
            linear_acceleration_x, linear_acceleration_y, linear_acceleration_z = af_l[0], af_l[1], af_l[2]

            s_x = s_x + linear_acceleration_x * 0.01
            s_y = s_y + linear_acceleration_y * 0.01
            s_z = s_z + (linear_acceleration_z - 9.8) * 0.01

            angular_velocity_x, angular_velocity_y, angular_velocity_z = wf_l[0], wf_l[1], wf_l[2]

            roll, pitch, yaw = ef_l[0], ef_l[1], ef_l[2]
            print("--- angle ---")
            heading = math.atan2(pitch, roll) + declination
            if(heading > 2*pi):
                heading = heading - 2*pi
            if(heading < 0):
                heading = heading + 2*pi
                
            print(roll,pitch,yaw)
            heading_angle = int(heading * 180/pi)
            print ("Heading Angle = %d°" %heading_angle)         
           

  
