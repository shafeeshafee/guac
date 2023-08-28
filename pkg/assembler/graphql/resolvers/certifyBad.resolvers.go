package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestCertifyBad is the resolver for the ingestCertifyBad field.
func (r *mutationResolver) IngestCertifyBad(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType model.MatchFlags, certifyBad model.CertifyBadInputSpec) (string, error) {
	ingestedCertifyBad, err := r.Backend.IngestCertifyBad(ctx, subject, &pkgMatchType, certifyBad)
	if err != nil {
		return "", err
	}
	return ingestedCertifyBad.ID, err
}

// IngestCertifyBads is the resolver for the ingestCertifyBads field.
func (r *mutationResolver) IngestCertifyBads(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType model.MatchFlags, certifyBads []*model.CertifyBadInputSpec) ([]string, error) {
	ingestedCertifyBads, err := r.Backend.IngestCertifyBads(ctx, subjects, &pkgMatchType, certifyBads)
	ingestedCertifyBadsIDS := []string{}
	if err == nil {
		for _, certifybad := range ingestedCertifyBads {
			ingestedCertifyBadsIDS = append(ingestedCertifyBadsIDS, certifybad.ID)
		}
	}
	return ingestedCertifyBadsIDS, err
}

// CertifyBad is the resolver for the CertifyBad field.
func (r *queryResolver) CertifyBad(ctx context.Context, certifyBadSpec model.CertifyBadSpec) ([]*model.CertifyBad, error) {
	return r.Backend.CertifyBad(ctx, &certifyBadSpec)
}
