# How to setup env:
# install cuda 12.1 and cudnn(add env vars)
# check cuda version: nvcc -V
# conda install python
# pip3 install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu121
# pip install ultralytics
# pip install lapx

import torch, cv2
from ultralytics import YOLO
from ultralytics.solutions import speed_estimation, distance_calculation

print(torch.cuda.is_available())

# Load the YOLOv8 model
model = YOLO('yolov8n.pt')

# Open the video file
cap = cv2.VideoCapture(0)

# Init speed-estimation obj
line_pts = [[0, 360], [1280, 360]]
speed_obj = speed_estimation.SpeedEstimator()
speed_obj.set_args(reg_pts=line_pts, names=model.model.names)

# Init dist obj
dist_obj = distance_calculation.DistanceCalculation()
dist_obj.set_args(names=model.model.names)

boxesSizes = {}
sizeChangeThreshold = 30000

# Loop through the video frames
while cap.isOpened():
    # Read a frame from the video
    success, frame = cap.read()

    if success:
        # Run YOLOv8 tracking on the frame, persisting tracks between frames
        tracks = model.track(frame, persist=True, verbose=False, device="1")

        # Visualize the results on the frame
        annotated_frame = tracks[0].plot()
        for track in tracks:
            boxes = track.boxes.cpu().numpy()
            for i, label in enumerate(boxes.cls):
                if label == 0:
                    if (boxes.xywh is not None) and (boxes.id is not None) and (i < len(boxes.xywh)) and (i < len(boxes.id)):
                        boxSize = boxes.xywh[i][2] * boxes.xywh[i][3]
                        id = boxes.id[i]
                        if id in boxesSizes:
                            if boxSize - boxesSizes[id] > sizeChangeThreshold:
                                print("Alert! One person is approaching!")
                                print(boxSize)
                                print(boxesSizes[id])
                        boxesSizes[id] = boxSize

        frame_with_speed = speed_obj.estimate_speed(frame, tracks)

        frame_with_dist = dist_obj.start_process(frame, tracks)

        # Display the annotated frame
        # cv2.imshow("YOLOv8 Tracking", annotated_frame)
        # cv2.imshow("YOLOv8 Tracking with Dist", frame_with_dist)
        if frame_with_speed.size > 0:
            cv2.imshow("YOLOv8 Tracking with Speed", frame_with_speed)

        # Break the loop if 'q' is pressed
        if cv2.waitKey(1) & 0xFF == ord("q"):
            break
    else:
        # Break the loop if the end of the video is reached
        break

# Release the video capture object and close the display window
cap.release()
cv2.destroyAllWindows()
