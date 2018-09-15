package ppic

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// imageEncoder represents a function which can encode an image.
type imageEncoder func(io.Writer, image.Image) error

// getImageSize extracts an image size from a set of URL values.
func getImageSize(q url.Values) (int, error) {
	ss := q.Get("size")

	if len(ss) == 0 {
		return 512, nil
	}

	s, err := strconv.Atoi(ss)

	if err != nil {
		return 0, err
	}

	return s, nil
}

func getPalette(q url.Values) Palette {
	pn := q.Get("palette")
	pa := DefaultPalette

	// If there is a palette name then look it up.
	if len(pn) > 0 {
		lpn := strings.ToLower(pn)

		for n, p := range Palettes {
			if lpn == strings.ToLower(n) {
				pa = p
			}
		}
	}

	// If the request is to invert the palette then do so.
	if q["inverse"] != nil {
		return pa.Inverse()
	}

	return pa
}

// getImageEncoder returns an imageEncoder for the specified path.
func getImageEncoder(p string) imageEncoder {
	ext := path.Ext(p)

	switch strings.ToLower(ext) {
	case ".gif":
		return func(w io.Writer, i image.Image) error {
			return gif.Encode(w, i, &gif.Options{NumColors: 2})
		}
	case ".jpg", ".jpeg":
		return func(w io.Writer, i image.Image) error {
			return jpeg.Encode(w, i, &jpeg.Options{Quality: 1})
		}
	case "", ".png":
		return func(w io.Writer, i image.Image) error {
			enc := png.Encoder{CompressionLevel: png.NoCompression}

			return enc.Encode(w, i)
		}
	default:
		return nil
	}
}

// Handler serves HTTP requests with generated images.
func Handler(res http.ResponseWriter, req *http.Request) {
	// We only support GETing images.
	if req.Method != http.MethodGet {
		res.Header().Set("Allow", http.MethodGet)
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	encoder := getImageEncoder(req.URL.Path)

	// If we couldn't find an encoder then we couldn't understand the extension.
	if encoder == nil {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "error: unsupported file format")
		return
	}

	q := req.URL.Query()

	// Get the image size from the request.
	size, err := getImageSize(q)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "error: invalid size")
		return
	}

	// Get the color palette from the request.
	pal := getPalette(q)

	// Get the path without extension.
	txt := strings.TrimSuffix(req.URL.Path[1:], path.Ext(req.URL.Path))
	img, err := GenerateImage(txt, size, true, false, pal)

	// Check if an invalid size was specified.
	if err == ErrInvalidSize {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "error: %s", err)
		return
	}

	// Check if something else bad happened during generation.
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "error: %s", err)
		return
	}

	if err = encoder(res, img); err != nil {
		fmt.Fprintf(res, "error: %s", err)
	}
}
