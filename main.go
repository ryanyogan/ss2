package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/ryanyogan/shippy-service-vessel/proto/vessel"
)

// Repository defines the methods available for the Vessel
// service.
type Repository interface {
	FindAvailable(*vessel.Specification) (*vessel.Vessel, error)
}

// VesselRepository holds all the vessels ... what else can I say?
// given I am talking to myself
type VesselRepository struct {
	vessels []*vessel.Vessel
}

// FindAvailable - checks a Specification against a map of vessels,
// if the capacity and max weight are below a vessels capacity
// and max weith, then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *vessel.Specification) (*vessel.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("no vessel found by that specification")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *vessel.Specification, res *vessel.Response) error {
	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*vessel.Vessel{
		&vessel.Vessel{
			Id:        "vessel1001",
			Name:      "Bananana Boat",
			MaxWeight: 200000,
			Capacity:  500,
		},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	srv.Init()

	// Register our RPC implementation
	vessel.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
