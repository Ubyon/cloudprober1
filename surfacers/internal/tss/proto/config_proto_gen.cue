package proto

#SurfacerConf: {
	connectionString?:  string @protobuf(1,string,name=connection_string)
	topicName?:         string @protobuf(2,string,name=topic_name,#"default="cloudprober.metrics""#)
	isTls?:             bool   @protobuf(3,bool,name=is_tls,default)
	metricsBufferSize?: int64  @protobuf(10,int64,name=metrics_buffer_size,"default=10000")
}
