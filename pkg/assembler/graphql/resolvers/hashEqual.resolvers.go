package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestHashEqual is the resolver for the ingestHashEqual field.
func (r *mutationResolver) IngestHashEqual(ctx context.Context, artifact model.ArtifactInputSpec, otherArtifact model.ArtifactInputSpec, hashEqual model.HashEqualInputSpec) (string, error) {
	ingestedHashEqual, err := r.Backend.IngestHashEqual(ctx, artifact, otherArtifact, hashEqual)
	if err != nil {
		return "", err
	}
	return ingestedHashEqual.ID, err
}

// IngestHashEquals is the resolver for the ingestHashEquals field.
func (r *mutationResolver) IngestHashEquals(ctx context.Context, artifacts []*model.ArtifactInputSpec, otherArtifacts []*model.ArtifactInputSpec, hashEquals []*model.HashEqualInputSpec) ([]string, error) {
	funcName := "IngestHashEquals"
	ingestedHashEqualsIDS := []string{}
	if len(artifacts) != len(otherArtifacts) {
		return ingestedHashEqualsIDS, gqlerror.Errorf("%v :: uneven artifacts and other artifacts for ingestion", funcName)
	} else if len(artifacts) != len(hashEquals) {
		return ingestedHashEqualsIDS, gqlerror.Errorf("%v :: uneven artifacts and hashEquals for ingestion", funcName)
	}

	ingestedHashEquals, err := r.Backend.IngestHashEquals(ctx, artifacts, otherArtifacts, hashEquals)
	if err == nil {
		for _, hashEqual := range ingestedHashEquals {
			ingestedHashEqualsIDS = append(ingestedHashEqualsIDS, hashEqual.ID)
		}
	}
	return ingestedHashEqualsIDS, err
}

// HashEqual is the resolver for the HashEqual field.
func (r *queryResolver) HashEqual(ctx context.Context, hashEqualSpec model.HashEqualSpec) ([]*model.HashEqual, error) {
	if hashEqualSpec.Artifacts != nil && len(hashEqualSpec.Artifacts) > 2 {
		return nil, gqlerror.Errorf("HashEqual :: Provided spec has too many Artifacts")
	}
	return r.Backend.HashEqual(ctx, &hashEqualSpec)
}
