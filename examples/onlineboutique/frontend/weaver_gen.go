// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package frontend

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)
var _ codegen.LatestVersion = codegen.Version[[0][17]struct{}]("You used 'weaver generate' codegen version 0.17.0, but you built your code with an incompatible weaver module version. Try upgrading 'weaver generate' and re-running it.")

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(Server{}),
		Listeners: []string{"boutique"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		RefData: "⟦36ba6b75:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/productcatalogservice/T⟧\n⟦ad903f0a:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/currencyservice/T⟧\n⟦ae7426b7:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T⟧\n⟦3324d893:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/recommendationservice/T⟧\n⟦f76a2b4a:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/checkoutservice/T⟧\n⟦dd0dfbe8:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/shippingservice/T⟧\n⟦24712bd9:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/onlineboutique/adservice/T⟧\n⟦29a161ab:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→boutique⟧\n",
	})
}

// weaver.Instance checks.
var _ weaver.InstanceOf[weaver.Main] = (*Server)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*Server)(nil)

// Local stub implementations.

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

// Client stub implementations.

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

// Server stub implementations.

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}
