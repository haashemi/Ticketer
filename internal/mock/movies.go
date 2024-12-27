package mock

import (
	"context"
	_ "embed"
	"time"

	"github.com/haashemi/Ticketer/internal/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

var timezone = time.FixedZone("Tehran", 12600)

func MockMovies(ctx context.Context, db *postgres.Connection) error {
	for _, movie := range sampleMovies {
		if err := db.InsertMovie(ctx, movie); err != nil {
			return err
		}
	}

	return nil
}

var sampleMovies = []postgres.InsertMovieParams{
	{
		Name:             "Residency",
		Genres:           []string{"Documentary"},
		MovieTime:        65,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 8, 0, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "The Lost Of The Whale Shark",
		Genres:           []string{"Documentary"},
		MovieTime:        62,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 9, 5, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Bache Zerang",
		Genres:           []string{"Adventure", "Kids"},
		MovieTime:        85,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 10, 10, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Dadars",
		Genres:           []string{"Documentary"},
		MovieTime:        71,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 11, 45, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Hedgehog",
		Genres:           []string{"Comedy"},
		MovieTime:        99,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 12, 56, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Luca",
		Genres:           []string{"Kids"},
		MovieTime:        96,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 14, 37, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Screenshot",
		Genres:           []string{"Theater"},
		MovieTime:        60,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 16, 13, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Human-Horse, Fifty-Fifty",
		Genres:           []string{"Theater"},
		MovieTime:        75,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 17, 13, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Melancholy",
		Genres:           []string{"Theater"},
		MovieTime:        60,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 18, 30, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Captive Breeding",
		Genres:           []string{"Documentary"},
		MovieTime:        68,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 19, 30, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Parvande Tamarrode Man",
		Genres:           []string{"Documentary"},
		MovieTime:        74,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 20, 40, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
	{
		Name:             "Khallaj Family",
		Genres:           []string{"Documentary"},
		MovieTime:        97,
		PremiereTime:     pgtype.Timestamptz{Time: time.Date(0, 0, 0, 21, 55, 0, 0, timezone), Valid: true},
		PremiereFromDate: pgtype.Date{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
		PremiereToDate:   pgtype.Date{Time: time.Date(2031, 1, 1, 0, 0, 0, 0, timezone), Valid: true},
	},
}
