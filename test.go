package main

import (
	"fmt"
	"os"
  )


buffer, err := bimg.Read("testdata/test.jpg")
if err != nil {
  fmt.Fprintln(os.Stderr, err)
}

newImage, err := bimg.NewImage(buffer).Convert(bimg.AVIF)
if err != nil {
  fmt.Fprintln(os.Stderr, err)
}

if bimg.NewImage(newImage).Type() == "avif" {
  fmt.Fprintln(os.Stderr, "The image was converted into avif")
}

bimg.Write("/tmp/new.jpg", newImage)