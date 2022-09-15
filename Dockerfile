FROM swr.ap-southeast-3.myhuaweicloud.com/root-common/java-base:oracle-jdk-1.8.0_202
#从build_image.sh传进来工程目录
ARG PROJECT_NAME
# copy run jar file
COPY ./super-view-server /data/code/
COPY ./run.sh  /data/code/
RUN sudo chmod o+x /data/code/run.sh    && \
    sudo chown -R ${AppUser}:${AppGroup} /data

# set workdir
WORKDIR /data/code
CMD ["/bin/bash","-c","bash /data/code/run.sh $Env"]