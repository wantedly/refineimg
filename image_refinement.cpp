#include "image_refinement.h"
#include <cmath>
#include <algorithm>

using namespace std;

namespace imgutil {

// cf. http://docs.opencv.org/2.4/doc/tutorials/imgproc/histograms/histogram_calculation/histogram_calculation.html
cv::Mat refine(const cv::Mat& src, cv::Mat* hist_image) {
  const int pixels = src.rows * src.cols;
  vector<cv::Mat> bgr_planes;
  cv::split(src, bgr_planes);
  int histSize = 256;

  /// Set the ranges (for B,G,R))
  float range[] = { 0, 256 } ;
  const float* histRange = { range };

  bool uniform = true; bool accumulate = false;

  cv::Mat hist[3];
  int low01[3], high01[3];
  int high_bar = pixels / 100;
  int low_bar = pixels - high_bar;

  /// Compute the histograms:
  for (int c = 0; c < 3; c++) {
    cv::calcHist(&bgr_planes[c], 1, 0, cv::Mat(), hist[c], 1, &histSize, &histRange, uniform, accumulate);
    int cnt = 0;
    bool first = true;
    for (int i = histSize - 1; i >= 0; i--) {
      int v = hist[c].at<float>(i);
      if (cnt < high_bar && cnt+v >= high_bar) high01[c] = i;
      if (cnt < low_bar && cnt+v >= low_bar) low01[c] = i;
      cnt += v;
    }
  }

  int low = min(min(128, low01[0]), min(low01[1], low01[2]));
  int high = max(max(128, high01[0]), max(high01[1], high01[2]));
  vector<unsigned char> lut(256);
  int l = -4, h = 5;
  double ll = 1.0 / (1 + pow(2, -l));
  double hh = 1.0 / (1 + pow(2, -h));

  for (int i = 0; i < 256; i++) {
    double x = 1.0 * (i - low) / (high - low) * (h - l) + l;
    double y = (1.0 / (1.0 + pow(2.0, -x)) - ll) / (hh - ll);
    // cout << low << " " << high << " " << i << " " << x << " " << y << endl;
    lut[i] = max(0.0, min(255.0, 255.0 * y));
  }
  cv::Mat dst;
  cv::LUT(src, lut, dst);

  if (hist_image != NULL) {
    // Draw the histograms for B, G and R
    int hist_w = 512; int hist_h = 400;
    int bin_w = cvRound((double) hist_w/histSize);
    *hist_image = cv::Mat(hist_h, hist_w, CV_8UC3, cv::Scalar(0,0,0));

    /// Normalize the result to [ 0, hist_image->rows ]
    cv::Mat histf[3];
    for (int c = 0; c < 3; c++) {
      cv::normalize(hist[c], histf[c], 0, hist_image->rows, cv::NORM_MINMAX, -1, cv::Mat());
    }

    /// Draw for each channel
    cv::Scalar colors[] = {cv::Scalar(255, 0, 0), cv::Scalar(0, 255, 0), cv::Scalar(0, 0, 255)};
    for (int c = 0; c < 3; c++) {
      for (int i = 1; i < histSize; i++) {
        cv::line(*hist_image, cv::Point(bin_w*(i-1), hist_h - cvRound(histf[c].at<float>(i-1))),
                       cv::Point(bin_w*(i), hist_h - cvRound(histf[c].at<float>(i))),
                       colors[c], 2);
      }
      cv::circle(*hist_image, cv::Point(bin_w*low01[c], hist_h*0.3), 10, colors[c], -1);
      cv::circle(*hist_image, cv::Point(bin_w*high01[c], hist_h*0.3), 10, colors[c], 2);
    }
    for (int i = 1; i < histSize; i++) {
      cv::line(*hist_image, cv::Point(bin_w*(i-1), hist_h * (255 - lut[i-1]) / 255.0), cv::Point(bin_w*(i), hist_h * (255 - lut[i]) / 255.0), cv::Scalar(0, 255, 255), 1);
    }
  }
  return dst;
}

}
