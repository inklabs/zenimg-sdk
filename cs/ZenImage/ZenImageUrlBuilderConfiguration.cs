using System;
using System.Text.RegularExpressions;

namespace ZenImage
{
    public class ZenImageUrlBuilderConfiguration
    {
        private string _imageCode;
        private string _pixelSize;
        private string _actualSize;
        private int? _background;
        private int? _backgroundTexture;

        public string ZenImageServiceLocation { get { return "http://i.zenimg.com";  } }
        public string Style { get; set; }
        public bool UseImageWrap { get; set; }
        public string EdgeColor { get; set; }
        public double? WrapDepth { get; set; }
        public double? EdgeDepth { get; set; }
        public double? RoundedDepth { get; set; }
        public string FrameCode { get; set; }
        public bool UseShadow { get; set; }
        public int? Pan { get; set; }
        public int? Tilt { get; set; }
        public int? Roll { get; set; }
        public string BackgroundTextureColor { get; set; }
        
        /// <summary>
        /// Gets and sets the URL of the original image.
        /// </summary>
        public string ImageUrl { get; set; }

        /// <summary>
        /// Gets and sets the image format (e.g. jpg, png)
        /// </summary>
        public string ImageFormat { get; set; }

        /// <summary>
        /// Gets and sets the image code.
        /// </summary>
        public string ImageCode
        {
            get { return _imageCode; }
            set
            {
                if (!Regex.IsMatch(value, "[a-zA-Z0-9]."))
                {
                    throw new ArgumentException("ImageCode must be of format [a-zA-Z0-9].");
                }

                _imageCode = value;
            }
        }

        public string ActualSize
        {
            get { return _actualSize; }
            set { _actualSize = value.ToUpper(); }
        }

        public string PixelSize
        {
            get { return _pixelSize; }
            set { _pixelSize = value.ToUpper(); }
        }

        public int? Background
        {
            get { return _background; }
            set
            {
                if (value != null && !(0 <= value && value <= 9))
                {
                    throw new ArgumentOutOfRangeException("Background value must in the range of [0-9]", "value");
                }

                _background = value;
            }
        }

        public int? BackgroundTexture
        {
            get { return _backgroundTexture; }
            set
            {
                if (value != null && !(0 <= value && value <= 9))
                {
                    throw new ArgumentOutOfRangeException("BackgroundTexture value must in the range of [0-9]", "value");
                }

                _backgroundTexture = value;
            }
        }

        /// <summary>
        /// Creates a new ZenImage URL builder.
        /// </summary>
        /// <returns>a configured ZenImageUrlBuilder</returns>
        public ZenImageUrlBuilder CreateBuilder()
        {
            return new ZenImageUrlBuilder(this);
        }

        /// <summary>
        /// Determines if the minimum required configuration options have been set.
        /// </summary>
        /// <returns>true if the configuration is valid; false otherwise.</returns>
        public bool IsConfigurationValid()
        {
            return !string.IsNullOrEmpty(Style) &&
                   !string.IsNullOrEmpty(ImageFormat) &&
                   !string.IsNullOrEmpty(PixelSize) &&
                   (!string.IsNullOrEmpty(ImageUrl) || !string.IsNullOrEmpty(ImageCode));
        }
        
        public ZenImageUrlBuilderConfiguration SetStyle(string style)
        {
            this.Style = style;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetImageFormat(string imageFormat)
        {
            this.ImageFormat = imageFormat;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetImageUrl(string imageUrl)
        {
            this.ImageUrl = imageUrl;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetUseImageWrap(bool useImageWrap)
        {
            this.UseImageWrap = useImageWrap;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetWrapDepth(double? wrapDepth)
        {
            this.WrapDepth = wrapDepth;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetEdgeDepth(double? edgeDepth)
        {
            this.EdgeDepth = edgeDepth;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetEdgeColor(string edgeColor)
        {
            this.EdgeColor = edgeColor;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetRoundedDepth(double? roundedDepth)
        {
            this.RoundedDepth = roundedDepth;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetFrameCode(string frameCode)
        {
            this.FrameCode = frameCode;
            
            return this;
        }

        public ZenImageUrlBuilderConfiguration SetBackground(int? background)
        {
            this.Background = background;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetBackgroundTexture(int? backgroundTexture)
        {
            this.BackgroundTexture = backgroundTexture;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetBackgroundTextureColor(string backgroundTextureColor)
        {
            this.BackgroundTextureColor = backgroundTextureColor;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetUseShadow(bool useShadow)
        {
            this.UseShadow = useShadow;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetPan(int? pan)
        {
            this.Pan = pan;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetTilt(int? tilt)
        {
            this.Tilt = tilt;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetRoll(int? roll)
        {
            this.Roll = roll;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetActualSize(string actualSize)
        {
            this.ActualSize = actualSize;

            return this;
        }

        public ZenImageUrlBuilderConfiguration SetPixelSize(string size)
        {
            this.PixelSize = size;

            return this;
        }
    }
}
