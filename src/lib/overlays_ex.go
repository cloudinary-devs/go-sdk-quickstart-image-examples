package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Overlays(cld *cloudinary.Cloudinary) string {
    img_coffee, err := cld.Image("cld-sample")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    img_coffee.Transformation = "w_400,h_250,c_fill,g_south/l_cld-sample-2,w_1.3,h_1.3,g_faces,c_crop/e_saturation:50/e_vignette/g_north_east,fl_layer_apply,w_100,r_max,y_20,x_90/l_cld-sample-3,h_55/e_hue:-20/fl_layer_apply,g_south/l_text:Cookie_40_bold:Love,co_rgb:f08/a_20,fl_layer_apply,x_35,y_44,g_south_west/c_crop,w_300,h_250,x_30/r_60"
    // Generate and print the delivery URL
    url, err := img_coffee.String()
    fmt.Println(url)
    if err != nil {
        fmt.Println("error")
    }
    return url
}