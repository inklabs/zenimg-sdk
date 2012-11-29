using System;
using Machine.Specifications;

namespace ZenImage.Test
{   
    [Subject(typeof(ZenImageUrlBuilder))]
    public class ZenImageUrlBuilderSpec
    {
        public abstract class concern
        {
            Establish c = () =>
            {
                zenimage_configuration = ZenImageUrlBuilder.Configure()
                                                           .SetStyle(ZenImageOptions.Styles.CanvasGallaryWrap)
                                                           .SetImageFormat("jpg")
                                                           .SetImageUrl("http://testserver/testimage.jpg")
                                                           .SetPixelSize("200x200");
            };

            protected static ZenImageUrlBuilder zenimage_url_builder;
            protected static ZenImageUrlBuilderConfiguration zenimage_configuration;
            protected static string zenimage_url;
        }

        public class when_creating_a_url_builder_with_an_invalid_configuration : concern
        {
            private static Exception Exception;

            Because b = () =>
            {
                Exception = Catch.Exception(() => ZenImageUrlBuilder.Configure().CreateBuilder());
            };

            It should_throw_an_invalid_argument_exception = () =>
                Exception.ShouldBeOfType<ArgumentException>();
        }

        public class when_creating_a_zenimage_url : concern
        {
            private static string image_url;

            Establish c = () =>
            {
                image_url = "http://testserver/testimage.jpg";

                zenimage_url_builder = ZenImageUrlBuilder.Configure()
                                                         .SetStyle(ZenImageOptions.Styles.CanvasGallaryWrap)
                                                         .SetImageFormat("jpg")
                                                         .SetImageUrl(image_url)
                                                         .SetPixelSize("200x200")
                                                         .CreateBuilder();
            };

            Because b = () =>
            {
                zenimage_url = zenimage_url_builder.GetZenImageUrl();
            };

            It should_include_the_zenimage_server_location_in_the_url = () =>
            {
                zenimage_url.ShouldStartWith(zenimage_configuration.ZenImageServiceLocation);
            };

            It should_include_the_style_in_the_url = () =>
            {
                zenimage_url.ShouldContain(ZenImageOptions.Styles.CanvasGallaryWrap);
            };

            It should_include_the_image_format_in_the_url = () =>
            {
                zenimage_url.ShouldContain(".jpg");
            };

            It should_include_the_encoded_the_image_url = () =>
            {
                var encoded_image_url = Uri.EscapeDataString(image_url);

                zenimage_url.ShouldContain("?url=" + encoded_image_url);
            };
        }

        public class when_configuring_a_canvas_gallery_style : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.Style = ZenImageOptions.Styles.CanvasGallaryWrap;
            };

            public class when_the_canvas_gallery_image_wrap_option_is_selected : when_configuring_a_canvas_gallery_style
            {
                Establish c = () =>
                {
                    zenimage_configuration.UseImageWrap = true;
                };

                Because b = () =>
                {
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();
                };

                It should_include_the_image_wrap_option_in_the_url_query_string = () =>
                {
                    zenimage_url.ShouldContain("IW");
                };
            }

            public class when_an_edge_color_is_configured : when_configuring_a_canvas_gallery_style
            {
                Establish c = () =>
                {
                    zenimage_configuration.EdgeColor = "112233AA";
                };

                Because b = () =>
                {
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();
                };

                It should_include_the_edge_color_option_in_the_url_query_string = () =>
                {
                    zenimage_url.ShouldContain("EC" + zenimage_configuration.EdgeColor);
                };
            }

            public class when_a_wrap_depth_is_configured : when_configuring_a_canvas_gallery_style
            {
                Establish c = () =>
                {
                    zenimage_configuration.WrapDepth = 1.5;
                };

                Because b = () =>
                {
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();
                };

                It should_include_the_wrap_depth_option_in_the_url_query_string = () =>
                {
                    zenimage_url.ShouldContain("D" + zenimage_configuration.WrapDepth);
                };
            }
        }

        public class when_configuring_an_acrylic_style : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.Style = ZenImageOptions.Styles.Acrylic;
            };

            public class when_an_edge_depth_is_configured : when_configuring_an_acrylic_style
            {
                Establish c = () =>
                {
                    zenimage_configuration.EdgeDepth = 3.5;
                };

                Because b = () =>
                {
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();
                };

                It should_include_the_edge_depth_option_in_the_url = () =>
                {
                    zenimage_url.ShouldContain("ED" + zenimage_configuration.EdgeDepth);
                };
            }
        }

        public class when_configuring_an_aluminum_style : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.Style = ZenImageOptions.Styles.Aluminum;
            };

            public class when_an_edge_depth_is_configured : when_configuring_an_aluminum_style
            {
                Establish c = () =>
                {
                    zenimage_configuration.EdgeDepth = 2.5;
                };

                Because b = () =>
                {
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();
                };

                It should_include_the_edge_depth_option_in_the_url = () =>
                {
                    zenimage_url.ShouldContain("ED" + zenimage_configuration.EdgeDepth);
                };
            }

            public class when_a_rounded_depth_is_configured : when_configuring_an_aluminum_style
            {
                Establish c = () =>
                {
                    zenimage_configuration.RoundedDepth = 0.5;
                };

                Because b = () =>
                {
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();
                };

                It should_include_the_rounded_depth_option_in_the_url = () =>
                {
                    zenimage_url.ShouldContain("RD" + zenimage_configuration.RoundedDepth);
                };
            }
        }

        public class when_a_frame_code_is_configured : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.FrameCode = "xyz";
            };

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_frame_code_option_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_F{0}", zenimage_configuration.FrameCode));
        }

        public class when_a_background_is_configured : concern
        {
            public class when_the_background_is_valid : when_a_background_is_configured
            {
                Establish c = () =>
                {
                    zenimage_configuration.Background = 9;
                };

                Because b = () =>
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

                It should_include_the_background_option_in_the_url = () =>
                    zenimage_url.ShouldContain(string.Format("_BG{0}", zenimage_configuration.Background));
            }

            public class when_the_background_is_invalid : when_a_background_is_configured
            {
                private static Exception Exception;

                Because b = () =>
                    Exception = Catch.Exception(() => zenimage_configuration.Background = 10);

                It should_throw_an_argument_out_of_range_exception = () =>
                    Exception.ShouldBeOfType<ArgumentOutOfRangeException>();

            }
        }

        public class when_a_background_texture_is_configured : concern
        {
            public class when_the_background_texture_is_valid : when_a_background_texture_is_configured
            {
                Establish c = () =>
                {
                    zenimage_configuration.BackgroundTexture = 9;
                };

                Because b = () =>
                    zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

                It should_include_the_background_texture_option_in_the_url = () =>
                    zenimage_url.ShouldContain(string.Format("_BT{0}", zenimage_configuration.BackgroundTexture));
            }

            public class when_the_background_texture_is_invalid : when_a_background_texture_is_configured
            {
                private static Exception Exception;

                Because b = () =>
                    Exception = Catch.Exception(() => zenimage_configuration.BackgroundTexture = 10);

                It should_throw_an_argument_out_of_range_exception = () =>
                    Exception.ShouldBeOfType<ArgumentOutOfRangeException>();

            }
        }

        public class when_a_background_texture_color_is_configured : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.BackgroundTextureColor = "AA11BB22";
            };

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_background_texture_color_option_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_BTC{0}", zenimage_configuration.BackgroundTextureColor));
        }

        public class when_a_shadow_is_configured : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.UseShadow = true;
            };

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_shadow_option_in_the_url = () =>
                zenimage_url.ShouldContain("_SHD");

        }

        public class when_a_pan_value_is_configured : concern
        {
            Establish c = () =>
            {
                zenimage_configuration.Pan = 45;
            };

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_pan_option_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_P{0}", zenimage_configuration.Pan));

        }

        public class when_a_tilt_value_is_configured : concern
        {
            Establish c = () =>
                zenimage_configuration.Tilt = 35;

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_tilt_option_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_T{0}", zenimage_configuration.Tilt));
        }

        public class when_a_roll_value_is_configured : concern
        {
            Establish c = () =>
                zenimage_configuration.Roll = 25;

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_roll_option_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_R{0}", zenimage_configuration.Roll));

        }

        public class when_an_actual_size_value_is_configured : concern
        {
            Establish c = () =>
                zenimage_configuration.ActualSize = "200x200";

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_actual_size_option_in_uppercase_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_A{0}",
                                                                    zenimage_configuration.ActualSize.ToUpper()));
        }

        public class when_a_size_value_is_configured : concern
        {
            Establish c = () =>
                zenimage_configuration.PixelSize = "150x150";

            Because b = () =>
                zenimage_url = zenimage_configuration.CreateBuilder().GetZenImageUrl();

            It should_include_the_size_option_in_upper_case_in_the_url = () =>
                zenimage_url.ShouldContain(string.Format("_{0}", zenimage_configuration.PixelSize.ToUpper()));
        }

        public class when_using_the_fluent_configuration : concern
        {
            static string image_url;

            Because b = () =>
            {
                image_url = "http://pdf.buildasign.com/Proof.ashx?width=378&height=297&watermark=false&intent=Preview&tcid=614B5A59695550704137674F3446754F7839775775513D3D";

                zenimage_url = ZenImageUrlBuilder.Configure()
                                                 .SetImageUrl(image_url)
                                                 .SetImageFormat("jpg")
                                                 .SetStyle(ZenImageOptions.Styles.CanvasGallaryWrap)
                                                 .SetPixelSize("200x200")
                                                 .SetUseImageWrap(true)
                                                 .SetBackground(0)
                                                 .SetTilt(-20)
                                                 .SetPan(-25)
                                                 .SetEdgeColor("AA11BB22")
                                                 .SetWrapDepth(1.5)
                                                 .CreateBuilder()
                                                 .GetZenImageUrl();
            };

            It should_include_the_encoded_image_url = () =>
            {
                var encoded_image_url = Uri.EscapeDataString(image_url);

                zenimage_url.ShouldContain("?url=" + encoded_image_url);
            };
            
            It should_include_the_file_format = () =>
                zenimage_url.ShouldContain(".jpg");

            It should_include_the_style = () =>
                zenimage_url.ShouldContain(ZenImageOptions.Styles.CanvasGallaryWrap);

            It should_include_the_pixel_size_in_upper_case = () =>
                zenimage_url.ShouldContain("200X200");

            It should_include_the_image_wrap_option = () =>
                zenimage_url.ShouldContain("_IW");

            It should_include_the_background = () =>
                zenimage_url.ShouldContain("_BG0");

            It should_include_the_tilt = () =>
                zenimage_url.ShouldContain("_T-20");

            It should_include_the_pan = () =>
                zenimage_url.ShouldContain("_P-25");

            It should_include_the_edge_color = () =>
                zenimage_url.ShouldContain("_ECAA11BB22");

            It should_include_the_wrap_detph = () =>
                zenimage_url.ShouldContain("_D1.5");
        }
    }
}
