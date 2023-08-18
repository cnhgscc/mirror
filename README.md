# mirror

# protoc
# 编译的文件在 pkg/pb
make protoc

# 平滑重载
systemctl reload nginx
