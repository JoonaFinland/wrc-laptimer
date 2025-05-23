package events

import (
	"context"
	"log/slog"

	"github.com/majori/wrc-laptimer/internal/database"
	"github.com/majori/wrc-laptimer/pkg/telemetry"
)

func ProcessTelemetryEvents(ctx context.Context, db *database.Database, packetCh <-chan telemetry.TelemetryPacket) {
	for {
		select {
		case pkt := <-packetCh:
			switch pkt := pkt.(type) {
			case *telemetry.TelemetrySessionStart:
				err := db.FlushTelemetry()
				if err != nil {
					slog.Error("could not save telemetry", "error", err)
				}

				err = db.StartSession(pkt)
				if err != nil {
					slog.Error("could not save session", "error", err)
				}
				slog.Info("session started")

			case *telemetry.TelemetrySessionUpdate:
				err := db.AppendTelemetry(pkt)
				if err != nil {
					slog.Error("could not create new appender for telemetry", "error", err)
				}

			case *telemetry.TelemetrySessionEnd:
				err := db.FlushTelemetry()
				if err != nil {
					slog.Error("could not save telemetry", "error", err)
				}

				err = db.EndSession(pkt)
				if err != nil {
					slog.Error("could not end session", "error", err)
				}
				slog.Info("session ended")

			case *telemetry.TelemetrySessionPause:
				continue
			case *telemetry.TelemetrySessionResume:
				continue
			default:
				slog.Warn("unknown packet type", "type", pkt)
			}
		case <-ctx.Done():
			slog.Info("exiting...")
			return
		}
	}
}
