package main

import (
	"golang.org/x/tour/pic"
)

//http://tour.golang.org/moretypes/14

var (
	settings struct {
		Zoom uint64
		X float64
		Y float64
		CanvasWidth uint64
		CanvasHeight uint64
		MaxIterations uint8
		InvertEnabled bool
	}
	
	MAX_COLOR_VALUE uint8 = 255
	
	maxPixelIntensity uint8 = 0
	minPixelIntensity uint8 = settings.MaxIterations
)

func initializeSettings() {
	settings.Zoom = 1433600
	settings.X = 0.2575938088553293
	settings.Y = 0.001144005911690848
	settings.MaxIterations = MAX_COLOR_VALUE
	settings.InvertEnabled = false
}

func generatePixelIterationMap( x int, y int ) [][]uint8 {
	settings.CanvasWidth = uint64(x)
	settings.CanvasHeight = uint64(y)

	pixelIterationMap := initializePixelIterationMap()

	var xPixel, yPixel uint64
	for xPixel = 0; xPixel < settings.CanvasWidth; xPixel++ {
		x0 := mapXPixelToCartesian( xPixel )
		
		for yPixel = 0; yPixel < settings.CanvasHeight; yPixel++ {
			y0 := mapYPixelToCartesian( yPixel )
			
			var x, y float64 = 0, 0
			var iteration uint8 = 0

			for ( ( x*x - y*y < 4 ) && ( iteration < settings.MaxIterations ) ) {
				var xTemporary float64 = ( x*x - y*y + x0 )

				y = 2*x*y + y0

				x = xTemporary

				iteration++
			}
			
			pixelIterationMap[ xPixel ][ yPixel ] = iteration

			if iteration > maxPixelIntensity {
				maxPixelIntensity = iteration
			}
			if iteration < minPixelIntensity {
				minPixelIntensity = iteration
			}
		}
	}
	return pixelIterationMap
}

func initializePixelIterationMap() [][]uint8 {
	pixelIterationMap := make( [][]uint8, settings.CanvasWidth )
	pixelIterationMapColumns := make( []uint8, settings.CanvasWidth * settings.CanvasHeight )
	for i := range pixelIterationMap {
		pixelIterationMap[ i ], pixelIterationMapColumns = pixelIterationMapColumns[ :settings.CanvasHeight ], pixelIterationMapColumns[ settings.CanvasHeight: ]
	}
	return pixelIterationMap
}

func mapXPixelToCartesian( xPixel uint64 ) float64 {
	xCartesian := ( ( float64( xPixel ) - ( float64( settings.CanvasWidth ) / 2 ) ) / float64( settings.Zoom ) ) + float64( settings.X )
	return xCartesian
}

func mapYPixelToCartesian( yPixel uint64 ) float64 {
	yCartesian := ( -1 * ( ( float64( yPixel ) - ( float64( settings.CanvasHeight ) / 2 ) ) / float64( settings.Zoom ) ) ) + float64( settings.Y )
	return yCartesian
}

func getPixelIntensityColor( pixelIntensity uint8 ) uint8 {
	optimizedPixelIntensity := pixelIntensity
	
	if settings.InvertEnabled {
		optimizedPixelIntensity = ( MAX_COLOR_VALUE - optimizedPixelIntensity )
	}
	
	return optimizedPixelIntensity
}

func main() {
	initializeSettings()
	pic.Show(generatePixelIterationMap)
}