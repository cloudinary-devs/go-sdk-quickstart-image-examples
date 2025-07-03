package main

import (
    "log"
    "html/template"
    "net/http"
    "github.com/cloudinary/cloudinary-go"
    ec "go_image_examples/src/lib"
    "os"
)

type Transformation struct {
    Description string
    Asset string
    Link string
    Text string
    Done  bool
}

type TransformationPageData struct {
    PageTitle string
    Transformations     []Transformation
}

func main() {
    cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

    url_cartoon := ec.Cartoonify(cld) 
    url_resize := ec.Resize(cld)
    url_resize_crop := ec.Resize_crop(cld)
    url_overlays := ec.Overlays(cld) 
    url_faces_gravity := ec.Faces_gravity(cld)
    url_auto_format := ec.Auto_format(cld)
    url_specified_format := ec.Specified_format(cld)
    url_sepia := ec.Sepia(cld)

    tmpl := template.Must(template.ParseFiles("transformations.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TransformationPageData{
            PageTitle: "Cloudinary Transformations: Go SDK",
            Transformations: []Transformation{
                {Description: "Cartoonified", Asset: url_cartoon, Done: false, Text: "Applying image effects and filters", Link: "https://cloudinary.com/documentation/go_media_transformations#applying_image_effects_and_filters"},
                {Description: "Resized", Asset: url_resize, Done: true, Text: "Add Cloudinary to your code", Link: "https://cloudinary.com/documentation/go_quick_start#2_add_cloudinary_to_your_code"},
                {Description: "Resized & Cropped", Asset: url_resize_crop, Done: true, Text: "Resizing and Cropping", Link: "https://cloudinary.com/documentation/go_media_transformations#resizing_and_cropping"},
                {Description: "Text and Image Overlays", Asset: url_overlays, Done: true, Text: "Adding text and image overlays", Link: "https://cloudinary.com/documentation/go_media_transformations#adding_text_and_image_overlays"},
                {Description: "Gravity with Face Detection", Asset: url_faces_gravity, Done: true, Text: "Combining transformations", Link: "https://cloudinary.com/documentation/go_media_transformations#combining_transformations"},
                {Description: "Auto format", Asset: url_auto_format, Done: true, Text: "Image optimizations", Link: "https://cloudinary.com/documentation/go_media_transformations#image_optimizations"},
                {Description: "Specified format", Asset: url_specified_format, Done: true, Text: "Converting to another image format", Link: "https://cloudinary.com/documentation/go_media_transformations#converting_to_another_image_format"},
                {Description: "Sepia effect", Asset: url_sepia, Done: true, Text: "Return the delivery URL", Link: "https://cloudinary.com/documentation/go_media_transformations#3_return_the_delivery_url"},
            },
        }
        if err := tmpl.Execute(w, data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    log.Println("Listening on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

