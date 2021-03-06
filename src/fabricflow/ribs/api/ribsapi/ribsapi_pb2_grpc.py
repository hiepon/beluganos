# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import ribsapi_pb2 as ribsapi__pb2


class RIBSCoreApiStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ModRib = channel.unary_unary(
        '/ribsapi.RIBSCoreApi/ModRib',
        request_serializer=ribsapi__pb2.RibUpdate.SerializeToString,
        response_deserializer=ribsapi__pb2.ModRibReply.FromString,
        )
    self.MonitorRib = channel.unary_stream(
        '/ribsapi.RIBSCoreApi/MonitorRib',
        request_serializer=ribsapi__pb2.MonitorRibRequest.SerializeToString,
        response_deserializer=ribsapi__pb2.RibUpdate.FromString,
        )
    self.SyncRib = channel.unary_unary(
        '/ribsapi.RIBSCoreApi/SyncRib',
        request_serializer=ribsapi__pb2.SyncRibRequest.SerializeToString,
        response_deserializer=ribsapi__pb2.SyncRibReply.FromString,
        )


class RIBSCoreApiServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def ModRib(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def MonitorRib(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SyncRib(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RIBSCoreApiServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ModRib': grpc.unary_unary_rpc_method_handler(
          servicer.ModRib,
          request_deserializer=ribsapi__pb2.RibUpdate.FromString,
          response_serializer=ribsapi__pb2.ModRibReply.SerializeToString,
      ),
      'MonitorRib': grpc.unary_stream_rpc_method_handler(
          servicer.MonitorRib,
          request_deserializer=ribsapi__pb2.MonitorRibRequest.FromString,
          response_serializer=ribsapi__pb2.RibUpdate.SerializeToString,
      ),
      'SyncRib': grpc.unary_unary_rpc_method_handler(
          servicer.SyncRib,
          request_deserializer=ribsapi__pb2.SyncRibRequest.FromString,
          response_serializer=ribsapi__pb2.SyncRibReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'ribsapi.RIBSCoreApi', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class RIBSApiStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetNexthops = channel.unary_stream(
        '/ribsapi.RIBSApi/GetNexthops',
        request_serializer=ribsapi__pb2.GetNexthopsRequest.SerializeToString,
        response_deserializer=ribsapi__pb2.Nexthop.FromString,
        )
    self.GetRics = channel.unary_stream(
        '/ribsapi.RIBSApi/GetRics',
        request_serializer=ribsapi__pb2.GetRicsRequest.SerializeToString,
        response_deserializer=ribsapi__pb2.RicEntry.FromString,
        )
    self.GetNexthopMap = channel.unary_stream(
        '/ribsapi.RIBSApi/GetNexthopMap',
        request_serializer=ribsapi__pb2.GetIPMapRequest.SerializeToString,
        response_deserializer=ribsapi__pb2.IPMapEntry.FromString,
        )


class RIBSApiServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetNexthops(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetRics(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetNexthopMap(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RIBSApiServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetNexthops': grpc.unary_stream_rpc_method_handler(
          servicer.GetNexthops,
          request_deserializer=ribsapi__pb2.GetNexthopsRequest.FromString,
          response_serializer=ribsapi__pb2.Nexthop.SerializeToString,
      ),
      'GetRics': grpc.unary_stream_rpc_method_handler(
          servicer.GetRics,
          request_deserializer=ribsapi__pb2.GetRicsRequest.FromString,
          response_serializer=ribsapi__pb2.RicEntry.SerializeToString,
      ),
      'GetNexthopMap': grpc.unary_stream_rpc_method_handler(
          servicer.GetNexthopMap,
          request_deserializer=ribsapi__pb2.GetIPMapRequest.FromString,
          response_serializer=ribsapi__pb2.IPMapEntry.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'ribsapi.RIBSApi', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
