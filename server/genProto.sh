# Color
RED=$(echo  '\033[00;31m')
GREEN=$(echo  '\033[00;32m')
YELLOW=$(echo '\033[00;33m')
BLUE=$(echo  '\033[00;34m')
MAGENTA=$(echo  '\033[00;35m')
PURPLE=$(echo '\033[00;35m')
CYAN=$(echo '\033[00;36m')
LIGHTGRAY=$(echo  '\033[00;37m')
LRED=$(echo  '\033[01;31m')
LGREEN=$(echo  '\033[01;32m')
LYELLOW=$(echo '\033[01;33m')
LBLUE=$(echo  '\033[01;34m')
LMAGENTA=$(echo  '\033[01;35m')
LPURPLE=$(echo  '\033[01;35m')
LCYAN=$(echo '\033[01;36m')
WHITE=$(echo  '\033[01;37m')
END=$(echo '\033[0m')
 
 
 
 

#def PATH
PROTO_PATH=./auth/api
GO_OUT_PATH=./auth/api/gen/v1
# make PATH
echo $LGREEN"-> 创建PATH路径"$END $LCYAN$PROTO_PATH$END
mkdir -p $PROTO_PATH
echo $LGREEN"-> 创建PATH路径"$END $LCYAN$GO_OUT_PATH$END
mkdir -p $GO_OUT_PATH

# go
echo $LGREEN"-> 生成auth.pb.go"$END 
protoc -I=$PROTO_PATH --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH $PROTO_PATH/auth.proto
echo $LGREEN"-> 生成auth.pb.gw.go"$END
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/auth.yaml:$GO_OUT_PATH $PROTO_PATH/auth.proto

#js
PBTS_BIN_DIR=../wx/miniprogram/node_modules/.bin
PBTS_OUT_DIR=../wx/miniprogram/service/proto_gen/auth
echo $LGREEN"-> 创建JS 输出PATH路径"$END $LCYAN$PBTS_OUT_DIR$END
mkdir -p $PBTS_OUT_DIR

echo $LGREEN"-> 生成auth_pb.js"$END
$PBTS_BIN_DIR/pbjs -t static -w es6 $PROTO_PATH/auth.proto --nocreate --no-encode --no-dencode --no-verify --no-delimited -o $PBTS_OUT_DIR/auth_pb_tmp.js
echo 'import * as $protobuf from "protobufjs";\n' > $PBTS_OUT_DIR/auth_pb.js
cat $PBTS_OUT_DIR/auth_pb_tmp.js >> $PBTS_OUT_DIR/auth_pb.js
rm $PBTS_OUT_DIR/auth_pb_tmp.js
echo $LGREEN"-> 生成auth_pb.d.ts"$END
$PBTS_BIN_DIR/pbts -o $PBTS_OUT_DIR/auth_pb.d.ts $PBTS_OUT_DIR/auth_pb.js
echo $LGREEN"-> 完成"$END