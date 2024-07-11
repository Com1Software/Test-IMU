
# Assuming you have magnetometer readings mag_x, mag_y, and mag_z
mag_x, mag_y, mag_z = read_magnetometer()

# Calculate the heading
heading = math.atan2(mag_y, mag_x)

# Convert the heading from radians to degrees
heading_degrees = heading * (180.0 / math.pi)

# If the heading is negative, convert it to a value between 0 and 360
if heading_degrees < 0:
    heading_degrees += 360.0

print("Heading: ", heading_degrees)
