package services

import (
	"sl-api/controllers"
	"sl-api/models"
	"sl-api/services"
)

func InsertSample() (models.Sample, error) {
	sample := models.Sample{}

	sample.ReviewStateID = 1 /* Initial state (Unrevised) */

	if services.Database.Create(sample).Error != nil {
		return sample, services.BadParamsError
	}

	return sample, nil
}

func AssociateLiquids(sampleID uint, lAmounts controllers.SampleLiquidsInput) error {
	/* Insert Sample liquid */
	// Water
	wErr := InsertSampleLiquid(&(models.SampleLiquid{Quantity: lAmounts.WaterAmount, SampleID: sampleID}))
	if wErr != nil {
		return wErr
	}

	// Cooking Oil
	oErr := InsertSampleLiquid(&(models.SampleLiquid{Quantity: lAmounts.OilAmount, SampleID: sampleID}))
	if oErr != nil {
		return oErr
	}

	// Car Oil
	coErr := InsertSampleLiquid(&(models.SampleLiquid{Quantity: lAmounts.CarOilAmount, SampleID: sampleID}))
	if coErr != nil {
		return coErr
	}

	// Others
	otErr := InsertSampleLiquid(&(models.SampleLiquid{Quantity: lAmounts.OthersAmount, SampleID: sampleID}))
	if otErr != nil {
		return otErr
	}

	return nil
}

func AssociateSensorReadings(sampleID uint, sReadings controllers.SampleReadingsInput) error {
	// Liquid Temperature
	ltErr := insertSensorReading(controllers.TemperatureSensorType, sampleID, sReadings.LiquidTemperature)
	if ltErr != nil {
		return ltErr
	}

	// External temperature
	etErr := insertSensorReading(controllers.ExternalTemperatureSensorType, sampleID, sReadings.ExternalTemperature)
	if etErr != nil {
		return etErr
	}

	// Liquid level
	llErr := insertSensorReading(controllers.LiquidLevelSensorType, sampleID, sReadings.LiquidLevel)
	if llErr != nil {
		return llErr
	}

	// Turbidity
	tdErr := insertSensorReading(controllers.TurbiditySensorType, sampleID, sReadings.Turbidity)
	if tdErr != nil {
		return tdErr
	}

	// Humidity
	hErr := insertSensorReading(controllers.HumiditySensorType, sampleID, sReadings.Humidity)
	if hErr != nil {
		return hErr
	}

	// Water flow
	wfErr := insertSensorReading(controllers.WaterFlowSensorType, sampleID, sReadings.WaterFlow)
	if wfErr != nil {
		return wfErr
	}

	return nil
}

func GetSample(id string) (models.Sample, error) {
	sample := models.Sample{}

	services.Database.First(&sample, id)

	if sample.ID == 0 {
		return sample, services.NotFoundError
	}

	return sample, nil
}

func GetAllSamples() ([]models.Sample, error) {
	var allSamples []models.Sample

	services.Database.Find(&allSamples)

	if len(allSamples) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSamples, services.EmptyDbError
	}

	return allSamples, nil
}

func DeleteSample(id string) error {
	return services.Database.Delete(&models.Sample{}, id).Error
}
