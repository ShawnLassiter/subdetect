package main

import (
	"fmt"
	"image/color"
	"log"

	"unicode/utf8"

	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

func main() {
	// Open video file with GoCV
	videoPath := "Tron.Ares.mkv"
	vc, err := gocv.VideoCaptureFile(videoPath)
	if err != nil {
		log.Fatalf("Error opening video capture file: %s", err)
	}
	defer vc.Close()

	// Create a window to display the video
	window := gocv.NewWindow("OCR Bounding Boxes")
	defer window.Close()

	// Create a gocv.Mat to hold each frame
	img := gocv.NewMat()
	defer img.Close()

	// Create a gosseract client
	client := gosseract.NewClient()
	client.SetLanguage("fra")
	defer client.Close()

	fmt.Println("Starting video processing...")
	for {
		if ok := vc.Read(&img); !ok {
			fmt.Println("Finished processing video.")
			break
		}
		if img.Empty() {
			continue
		}

		// Encode the gocv.Mat to a JPEG byte buffer in memory
		buf, err := gocv.IMEncode(gocv.JPEGFileExt, img)
		if err != nil {
			log.Fatalf("Failed to encode frame to JPEG: %v", err)
		}

		// Pass the encoded JPEG bytes to gosseract
		err = client.SetImageFromBytes(buf.GetBytes())
		// It's important to release the buffer C-memory after use
		buf.Close()
		if err != nil {
			log.Fatalf("Could not set image from bytes: %v", err)
		}

		// Get bounding boxes for each word
		boxes, err := client.GetBoundingBoxes(gosseract.RIL_WORD)
		if err != nil {
			log.Fatalf("Could not get bounding boxes: %v", err)
		}

		// Draw each bounding box on the frame
		for _, box := range boxes {
			if box.Confidence < 50 || utf8.RuneCountInString(box.Word) < 2 {
				continue
			}
			// The rectangle to draw, from the BoundingBox struct
			rectangle := box.Box
			// The detected word
			word := box.Word

			fmt.Printf("Found word(%f): '%s' at %v\n", box.Confidence, word, rectangle)

			// Draw the rectangle on the original image (img)
			gocv.Rectangle(&img, rectangle, color.RGBA{R: 0, G: 255, B: 0, A: 0}, 2)
		}

		// Display the image with the bounding boxes in the window
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
