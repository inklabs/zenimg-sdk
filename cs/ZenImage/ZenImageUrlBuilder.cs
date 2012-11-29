using System;
using System.Collections.Generic;

namespace ZenImage
{
    /// <summary>
    /// This is a port of the ZenImage JavaScript API. This class receives a configuration and generates a URL string 
    /// that allows clients to retrieve images from ZenImage based on the configuration options embedded in the URL.
    /// 
    /// See http://www.zenimg.com/api for more information. 
    /// </summary>
    public class ZenImageUrlBuilder
    {
        private readonly ZenImageUrlBuilderConfiguration _configuration;

        public ZenImageUrlBuilder(ZenImageUrlBuilderConfiguration configuration)
        {
            if ( !configuration.IsConfigurationValid() )
            {
                throw new ArgumentException("The given configuration is not valid", "configuration");    
            }

            _configuration = configuration;
        }

        /// <summary>
        /// Creates a new ZenImage configuration for fluent configuration.
        /// </summary>
        /// <returns>a new ZenImage configuration</returns>
        public static ZenImageUrlBuilderConfiguration Configure()
        {
            return new ZenImageUrlBuilderConfiguration();
        }

        /// <summary>
        /// Creates a ZenImage URL based on the given configuration. 
        /// </summary>
        /// <param name="configuration">a configuration specifying the Z</param>
        /// <returns>a ZenImage URL.</returns>
        public static string GetZenImageUri(ZenImageUrlBuilderConfiguration configuration)
        {
            return new ZenImageUrlBuilder(configuration).GetZenImageUrl();
        }

        /// <summary>
        /// Creates a ZenImage URL based on the given configuration. 
        /// </summary>
        /// <returns>a ZenImage URL.</returns>
        public string GetZenImageUrl()
        {
            var options = new List<string> {_configuration.Style};

            options.AddRange( this.GetStyleSpecificOptions() );

            // Intentionally excluding the "Wood" style option from the configuration
            // due to some inconcistencies between the API documentation (http://www.zenimg.com/api)
            // and the actual code examples. The documentation defines the wood style code as "WD", 
            // whereas the code examples use "W". The code examples also allow for the useage of the 
            // "Wood" style in addition to other styles (it's outside of the style-related if-else-if block).
            // May need to check with the ZenImage folks for clarification.

            this.AddOption(ZenImageOptions.FrameCode, _configuration.FrameCode, options);
            this.AddOption(ZenImageOptions.Background, _configuration.Background, options);
            this.AddOption(ZenImageOptions.BackgroundTexture, _configuration.BackgroundTexture, options);
            this.AddOption(ZenImageOptions.BackgroundTextureColor, _configuration.BackgroundTextureColor, options);
            this.AddOption(ZenImageOptions.Shadow, _configuration.UseShadow, options);
            this.AddOption(ZenImageOptions.Pan, _configuration.Pan, options);
            this.AddOption(ZenImageOptions.Tilt, _configuration.Tilt, options);
            this.AddOption(ZenImageOptions.Roll, _configuration.Roll, options);
            this.AddOption(ZenImageOptions.ActualSize, _configuration.ActualSize, options);
            this.AddOption(ZenImageOptions.Size, _configuration.PixelSize, options);

            return BuildUrl(options);
        }

        /// <summary>
        /// Gets all style-related options.
        /// </summary>
        /// <returns>a list of style options to be added to the ZenImage URL</returns>
        private IEnumerable<string> GetStyleSpecificOptions()
        {
            if (_configuration.Style == ZenImageOptions.Styles.CanvasGallaryWrap)
            {
                return GetCanvasGalleryStyleOptions();
            }

            if (_configuration.Style == ZenImageOptions.Styles.Aluminum)
            {
                return GetAluminumStyleOptions();
            }

            if (_configuration.Style == ZenImageOptions.Styles.Acrylic)
            {
                return GetAcrylicStyleOptions();
            }
            
            throw new Exception("The configured ZenImage style is not supported.");
        }

        /// <summary>
        /// Gets all canvas gallery (CG) style options.
        /// </summary>
        /// <returns>a list of style options to be added to the ZenImage URL</returns>
        private IEnumerable<string> GetCanvasGalleryStyleOptions()
        {
            var options = new List<string>();

            this.AddOption(ZenImageOptions.ImageWrap, _configuration.UseImageWrap, options);
            this.AddOption(ZenImageOptions.EdgeColor, _configuration.EdgeColor, options);
            this.AddOption(ZenImageOptions.WrapDepth, _configuration.WrapDepth, options);

            return options;
        }

        /// <summary>
        /// Gets all aluminium (AL) style options.
        /// </summary>
        /// <returns>a list of style options to be added to the ZenImage URL</returns>
        private IEnumerable<string> GetAluminumStyleOptions()
        {
            var options = new List<string>();

            this.AddOption(ZenImageOptions.EdgeDepth, _configuration.EdgeDepth, options);
            this.AddOption(ZenImageOptions.RoundedDepth, _configuration.RoundedDepth, options);

            return options;
        }

        /// <summary>
        /// Gets all acrylic (AC) style options.
        /// </summary>
        /// <returns>a list of style options to be added to the ZenImage URL</returns>
        private IEnumerable<string> GetAcrylicStyleOptions()
        {
            var options = new List<string>();

            this.AddOption(ZenImageOptions.EdgeDepth, _configuration.EdgeDepth, options);

            return options;
        }

        /// <summary>
        /// Adds the provided value and prefix to the list of options as long as the value is not null.
        /// </summary>
        /// <param name="optionPrefix">a ZenImage option prefix</param>
        /// <param name="value">a ZenImage option value</param>
        /// <param name="options">a list of ZenImage options</param>
        private void AddOption(string optionPrefix, int? value, List<string> options)
        {
            if (value.HasValue)
            {
                options.Add(optionPrefix + value.Value);
            }    
        }

        /// <summary>
        /// Adds the provided value and prefix to the list of options as long as the value is not null.
        /// </summary>
        /// <param name="optionPrefix">a ZenImage option prefix</param>
        /// <param name="value">a ZenImage option value</param>
        /// <param name="options">a list of ZenImage options</param>
        private void AddOption(string optionPrefix, double? value, List<string> options)
        {
            if (value.HasValue)
            {
                options.Add(optionPrefix + value.Value);
            }    
        }

        /// <summary>
        /// Adds the provided value and prefix to the list of options as long as the value is not null or empty.
        /// </summary>
        /// <param name="optionPrefix">a ZenImage option prefix</param>
        /// <param name="value">a ZenImage option value</param>
        /// <param name="options">a list of ZenImage options</param>
        private void AddOption(string optionPrefix, string value, List<string> options)
        {
            if (!string.IsNullOrEmpty(value))
            {
                options.Add(optionPrefix + value);
            }
        }

        /// <summary>
        /// Adds the provided prefix to the list of options as long as value is true.
        /// </summary>
        /// <param name="optionPrefix">a ZenImage option prefix</param>
        /// <param name="value">indicates whether a ZenImage option is to be used</param>
        /// <param name="options">a list of ZenImage options</param>
        private void AddOption(string optionPrefix, bool value, List<string> options)
        {
            if (value)
            {
                options.Add(optionPrefix);
            }
        }

        /// <summary>
        /// Builds a string from the list of ZenImage options.
        /// </summary>
        /// <param name="options">a list of ZenImage options</param>
        /// <returns>an underscore-delimited string containing ZenImage options</returns>
        private string GetOptionsString(IEnumerable<string> options)
        {
            return string.Join("_", options) + "." + _configuration.ImageFormat;
        }

        /// <summary>
        /// Builds the ZenImage URL based on the list of options.
        /// </summary>
        /// <param name="options">a list of ZenImage options</param>
        /// <returns>the URL to the ZenImage image.</returns>
        private string BuildUrl(IEnumerable<string> options)
        {
            string zenImageUrl = null;

            var joinedOptions = GetOptionsString(options);

            if (!string.IsNullOrEmpty(_configuration.ImageCode))
            {
                zenImageUrl = string.Format("{0}/v1/{1}_{2}", _configuration.ZenImageServiceLocation, 
                                                              _configuration.ImageCode,
                                                              joinedOptions);
            } 
            else
            {
                var encodedImageUrl = Uri.EscapeDataString(_configuration.ImageUrl);

                zenImageUrl = string.Format("{0}/v1/url/{1}?url={2}", _configuration.ZenImageServiceLocation,
                                                                      joinedOptions, 
                                                                      encodedImageUrl);
            }

            return zenImageUrl;
        }


    }
}
