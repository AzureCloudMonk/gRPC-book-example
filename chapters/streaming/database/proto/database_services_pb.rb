# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/database.proto for package 'practical.grpc.v1'

require 'grpc'
require 'proto/database_pb'

module Practical
  module Grpc
    module V1
      module Database
        class Service

          include GRPC::GenericService

          self.marshal_class_method = :encode
          self.unmarshal_class_method = :decode
          self.service_name = 'practical.grpc.v1.Database'

          rpc :Search, SearchRequest, stream(SearchResponse)
        end

        Stub = Service.rpc_stub_class
      end
    end
  end
end
