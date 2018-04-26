# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/sfapi.proto for package ''

require 'grpc'
require 'proto/sfapi_pb'

module Starfriends
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'Starfriends'

    # Get a single Film by unique ID
    rpc :GetFilm, GetFilmRequest, GetFilmResponse
    # Get a list of all Films
    rpc :ListFilms, ListFilmsRequest, ListFilmsResponse
  end

  Stub = Service.rpc_stub_class
end
