# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from core.v1alpha1 import api_pb2 as core_dot_v1alpha1_dot_api__pb2


class EngineServiceStub(object):
    """**
    Service definition

    EngineService is the service that provides the core functionality of the engine and to access low-level operations
    over feature values.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FeatureDescriptor = channel.unary_unary(
                '/core.v1alpha1.EngineService/FeatureDescriptor',
                request_serializer=core_dot_v1alpha1_dot_api__pb2.FeatureDescriptorRequest.SerializeToString,
                response_deserializer=core_dot_v1alpha1_dot_api__pb2.FeatureDescriptorResponse.FromString,
                )
        self.Get = channel.unary_unary(
                '/core.v1alpha1.EngineService/Get',
                request_serializer=core_dot_v1alpha1_dot_api__pb2.GetRequest.SerializeToString,
                response_deserializer=core_dot_v1alpha1_dot_api__pb2.GetResponse.FromString,
                )
        self.Set = channel.unary_unary(
                '/core.v1alpha1.EngineService/Set',
                request_serializer=core_dot_v1alpha1_dot_api__pb2.SetRequest.SerializeToString,
                response_deserializer=core_dot_v1alpha1_dot_api__pb2.SetResponse.FromString,
                )
        self.Append = channel.unary_unary(
                '/core.v1alpha1.EngineService/Append',
                request_serializer=core_dot_v1alpha1_dot_api__pb2.AppendRequest.SerializeToString,
                response_deserializer=core_dot_v1alpha1_dot_api__pb2.AppendResponse.FromString,
                )
        self.Incr = channel.unary_unary(
                '/core.v1alpha1.EngineService/Incr',
                request_serializer=core_dot_v1alpha1_dot_api__pb2.IncrRequest.SerializeToString,
                response_deserializer=core_dot_v1alpha1_dot_api__pb2.IncrResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/core.v1alpha1.EngineService/Update',
                request_serializer=core_dot_v1alpha1_dot_api__pb2.UpdateRequest.SerializeToString,
                response_deserializer=core_dot_v1alpha1_dot_api__pb2.UpdateResponse.FromString,
                )


class EngineServiceServicer(object):
    """**
    Service definition

    EngineService is the service that provides the core functionality of the engine and to access low-level operations
    over feature values.
    """

    def FeatureDescriptor(self, request, context):
        """FeatureDescriptor returns the feature descriptor for the given selector.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Get returns the feature value or model prediction for the given selector.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Set(self, request, context):
        """Set sets the feature value for the given selector.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Append(self, request, context):
        """Append appends the given value to the feature value for the given selector.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Incr(self, request, context):
        """Incr increments the feature value for the given selector.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Update updates the feature value for the given selector.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EngineServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FeatureDescriptor': grpc.unary_unary_rpc_method_handler(
                    servicer.FeatureDescriptor,
                    request_deserializer=core_dot_v1alpha1_dot_api__pb2.FeatureDescriptorRequest.FromString,
                    response_serializer=core_dot_v1alpha1_dot_api__pb2.FeatureDescriptorResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=core_dot_v1alpha1_dot_api__pb2.GetRequest.FromString,
                    response_serializer=core_dot_v1alpha1_dot_api__pb2.GetResponse.SerializeToString,
            ),
            'Set': grpc.unary_unary_rpc_method_handler(
                    servicer.Set,
                    request_deserializer=core_dot_v1alpha1_dot_api__pb2.SetRequest.FromString,
                    response_serializer=core_dot_v1alpha1_dot_api__pb2.SetResponse.SerializeToString,
            ),
            'Append': grpc.unary_unary_rpc_method_handler(
                    servicer.Append,
                    request_deserializer=core_dot_v1alpha1_dot_api__pb2.AppendRequest.FromString,
                    response_serializer=core_dot_v1alpha1_dot_api__pb2.AppendResponse.SerializeToString,
            ),
            'Incr': grpc.unary_unary_rpc_method_handler(
                    servicer.Incr,
                    request_deserializer=core_dot_v1alpha1_dot_api__pb2.IncrRequest.FromString,
                    response_serializer=core_dot_v1alpha1_dot_api__pb2.IncrResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=core_dot_v1alpha1_dot_api__pb2.UpdateRequest.FromString,
                    response_serializer=core_dot_v1alpha1_dot_api__pb2.UpdateResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'core.v1alpha1.EngineService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class EngineService(object):
    """**
    Service definition

    EngineService is the service that provides the core functionality of the engine and to access low-level operations
    over feature values.
    """

    @staticmethod
    def FeatureDescriptor(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.v1alpha1.EngineService/FeatureDescriptor',
            core_dot_v1alpha1_dot_api__pb2.FeatureDescriptorRequest.SerializeToString,
            core_dot_v1alpha1_dot_api__pb2.FeatureDescriptorResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.v1alpha1.EngineService/Get',
            core_dot_v1alpha1_dot_api__pb2.GetRequest.SerializeToString,
            core_dot_v1alpha1_dot_api__pb2.GetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Set(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.v1alpha1.EngineService/Set',
            core_dot_v1alpha1_dot_api__pb2.SetRequest.SerializeToString,
            core_dot_v1alpha1_dot_api__pb2.SetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Append(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.v1alpha1.EngineService/Append',
            core_dot_v1alpha1_dot_api__pb2.AppendRequest.SerializeToString,
            core_dot_v1alpha1_dot_api__pb2.AppendResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Incr(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.v1alpha1.EngineService/Incr',
            core_dot_v1alpha1_dot_api__pb2.IncrRequest.SerializeToString,
            core_dot_v1alpha1_dot_api__pb2.IncrResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.v1alpha1.EngineService/Update',
            core_dot_v1alpha1_dot_api__pb2.UpdateRequest.SerializeToString,
            core_dot_v1alpha1_dot_api__pb2.UpdateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
