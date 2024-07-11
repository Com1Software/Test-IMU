
import serial
import time
import numpy as np

ser = serial.Serial()
ser.port = '/dev/ttyUSB0' 
ser.baudrate = 115200
ser.parity = 'N'
ser.bytesize = 8
ser.timeout = 1
ser.open()

readData = ""
dataStartRecording = False

print('Starting...', ser.name)
time.sleep(1)
ser.reset_input_buffer()

while True:
	rawData = ser.read(size=2).hex()
	if rawData == '5551':
		dataStartRecording = True
	if dataStartRecording == True:
		readData = readData + rawData
		if len(readData) == 88:
			StartAddress_3 = int(readData[44:46], 16)
			StartAddress_ypr = int(readData[46:48], 16)
			RollL = int(readData[48:50], 16)
			RollH = int(readData[50:52], 16)
			PitchL = int(readData[52:54], 16)
			PitchH = int(readData[54:56], 16)
			YawL = int(readData[56:58], 16)
			YawH = int(readData[58:60], 16)
			VL = int(readData[60:62], 16)
			VH = int(readData[62:64], 16)
			SUM_ypr = int(readData[64:66], 16)
			
						
			Roll = float(np.short((RollH<<8)|RollL)/32768.0*180.0)
			Pitch = float(np.short((PitchH<<8)|PitchL)/32768.0*180.0)
			Yaw = float(np.short((YawH<<8)|YawL)/32768.0*180.0)
			
			print("%7.3f" % Roll, "%7.3f" % Pitch, "%7.3f" % Yaw) # This maps out tilt angles of the axes.
			
			readData = ""
			dataStartRecording = False

	
