#Originally named Dockerfile-cpu, name it to Dockerfile
# Start FROM Ubuntu image https://hub.docker.com/_/ubuntu
FROM ubuntu:23.10

# Set environment variables
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_BREAK_SYSTEM_PACKAGES=1

# Downloads to user config dir
ADD https://github.com/ultralytics/assets/releases/download/v0.0.0/Arial.ttf \
    https://github.com/ultralytics/assets/releases/download/v0.0.0/Arial.Unicode.ttf \
    /root/.config/Ultralytics/

# Install linux packages
# g++ required to build 'tflite_support' and 'lap' packages, libusb-1.0-0 required for 'tflite_support' package
RUN apt update \
    && apt install --no-install-recommends -y python3-pip git zip unzip wget curl htop libgl1 libglib2.0-0 libpython3-dev gnupg g++ libusb-1.0-0 python3-venv && rm -rf /var/lib/apt/lists*
#added python3-venv and added line to remove cache dir

# Create working directory
WORKDIR /ultralytics

# Copy contents and configure git
COPY . .
RUN sed -i '/^\[http "https:\/\/github\.com\/"\]/,+1d' .git/config
ADD https://github.com/ultralytics/assets/releases/download/v8.2.0/yolov8n.pt .

# Install pip packages
RUN python3 -m pip install --upgrade pip wheel
RUN pip install -e ".[export]" --extra-index-url https://download.pytorch.org/whl/cpu

#Added additional libraries 
RUN pip install roboflow fastdup ultralytics opencv-python

#download colab py and txt files to cwd
RUN curl -L --remote-name-all https://raw.githubusercontent.com/YizhakTime/test_python/main/{pothole_detection.py,docker.txt}
#RUN curl -OL https://raw.githubusercontent.com/YizhakTime/test_python/main/pothole_detection.py 

# Run exports to AutoInstall packages
RUN yolo export model=tmp/yolov8n.pt format=edgetpu imgsz=32
RUN yolo export model=tmp/yolov8n.pt format=ncnn imgsz=32

#creates virtual python environment in cwd (i.e. ultralytics)
RUN python3 -m ./venv 

#enable venv
ENV PATH="venv/bin:$PATH"

#clone repo
RUN git clone git@github.com:roboflow-ai/auto-annotate.git

#install pkgs inside auto-annotate
RUN pip install -e ./auto-annotate

# Remove exported models
RUN rm -rf tmp

# Creates a symbolic link to make 'python' point to 'python3'
RUN ln -sf /usr/bin/python3 /usr/bin/python

# Usage Examples -------------------------------------------------------------------------------------------------------

# Build and Push
# t=ultralytics/ultralytics:latest-cpu && sudo docker build -f docker/Dockerfile-cpu -t $t . && sudo docker push $t

# Run
# t=ultralytics/ultralytics:latest-cpu && sudo docker run -it --ipc=host --name NAME $t

# Pull and Run
# t=ultralytics/ultralytics:latest-cpu && sudo docker pull $t && sudo docker run -it --ipc=host --name NAME $t

# Pull and Run with local volume mounted
# t=ultralytics/ultralytics:latest-cpu && sudo docker pull $t && sudo docker run -it --ipc=host -v "$(pwd)"/shared/datasets:/datasets $t
